#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D texture0;
uniform vec4 colDiffuse;

const vec2 size = vec2(1600, 900);   // Framebuffer size
const float samples = 13;            // Pixels per axis; higher = bigger glow, worse performance
const float quality = 0.8;           // Defines size factor: Lower = smaller glow, better quality

void main()
{
    // BLOOM
    vec4 sum = vec4(0);
    vec2 sizeFactor = vec2(1)/size*quality;

    // Texel color fetching from texture sampler
    vec4 source = texture(texture0, fragTexCoord);

    const int range = 3; // should be = (samples - 1)/2;

    for (int x = -range; x <= range; x++)
    {
        for (int y = -range; y <= range; y++)
        {
            sum += texture(texture0, fragTexCoord + vec2(x, y)*sizeFactor);
        }
    }

    // Calculate final fragment color
    vec4 bloomed = ((sum/(samples*samples)) + source)*colDiffuse;

    // CHROMATIC ABBERATION
    float r = texture(texture0, fragTexCoord + vec2(-0.0006, 0)).r;
    float b = texture(texture0, fragTexCoord + vec2(0.0006, 0)).b;

    vec4 chrabb = vec4(r, texture(texture0, fragTexCoord).g, b, fragColor.a) * colDiffuse;

    // COMBINE

    finalColor = bloomed/2+chrabb/2;
}
