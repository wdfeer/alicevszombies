#version 330 core

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D texture0;
uniform vec4 colDiffuse;

void main()
{
    const float offset = 0.0006;
    float r = texture(texture0, fragTexCoord + vec2(-offset, 0)).r;
    float b = texture(texture0, fragTexCoord + vec2(offset, 0)).b;

    finalColor = vec4(r, texture(texture0, fragTexCoord).g, b, fragColor.a) * colDiffuse;
}
