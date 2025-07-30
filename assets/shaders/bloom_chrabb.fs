#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D texture0;
uniform vec4 colDiffuse;

const vec2 size = vec2(1600, 900);
// Fewer samples and lower quality to compensate for only 50% bloom
const float samples = 9;
const float quality = 0.6;

void main()
{
    // BLOOM

    vec4 sum = vec4(0);
    vec2 sizeFactor = vec2(1)/size*quality;
    vec4 source = texture(texture0, fragTexCoord);
    const int range = 3; // bigger range to compensate for only 50% bloom
    for (int x = -range; x <= range; x++)
    {
        for (int y = -range; y <= range; y++)
        {
            sum += texture(texture0, fragTexCoord + vec2(x, y)*sizeFactor);
        }
    }
    vec4 bloom = ((sum/(samples*samples)) + source)*colDiffuse;

    // CHROMATIC ABBERATION

    const float offset = 0.0011; // bigger range to compensate for only 50% chromatic abberation
    float r = texture(texture0, fragTexCoord + vec2(-offset, 0)).r;
    float b = texture(texture0, fragTexCoord + vec2(offset, 0)).b;
    vec4 chrabb = vec4(r, texture(texture0, fragTexCoord).g, b, fragColor.a) * colDiffuse;

    // COMBINE

    finalColor = bloom/2+chrabb/2;
}
