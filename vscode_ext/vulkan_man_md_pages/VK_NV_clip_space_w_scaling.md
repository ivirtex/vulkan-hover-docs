# VK_NV_clip_space_w_scaling(3) Manual Page

## Name

VK_NV_clip_space_w_scaling - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

88

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Eric Werness <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_clip_space_w_scaling%5D%20@ewerness-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_clip_space_w_scaling%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ewerness-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-02-15

**Contributors**  
- Eric Werness, NVIDIA

- Kedarnath Thangudu, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Virtual Reality (VR) applications often involve a post-processing step
to apply a “barrel” distortion to the rendered image to correct the
“pincushion” distortion introduced by the optics in a VR device. The
barrel distorted image has lower resolution along the edges compared to
the center. Since the original image is rendered at high resolution,
which is uniform across the complete image, a lot of pixels towards the
edges do not make it to the final post-processed image.

This extension provides a mechanism to render VR scenes at a non-uniform
resolution, in particular a resolution that falls linearly from the
center towards the edges. This is achieved by scaling the w coordinate
of the vertices in the clip space before perspective divide. The clip
space w coordinate of the vertices **can** be offset as of a function of
x and y coordinates as follows:

w' = w + Ax + By

In the intended use case for viewport position scaling, an application
should use a set of four viewports, one for each of the four quadrants
of a Cartesian coordinate system. Each viewport is set to the dimension
of the image, but is scissored to the quadrant it represents. The
application should specify A and B coefficients of the w-scaling
equation above, that have the same value, but different signs, for each
of the viewports. The signs of A and B should match the signs of x and y
for the quadrant that they represent such that the value of w' will
always be greater than or equal to the original w value for the entire
image. Since the offset to w, (Ax + By), is always positive, and
increases with the absolute values of x and y, the effective resolution
will fall off linearly from the center of the image to its edges.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportWScalingNV.html)

- Extending
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html):

  - [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_CLIP_SPACE_W_SCALING_EXTENSION_NAME`

- `VK_NV_CLIP_SPACE_W_SCALING_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV`

## <a href="#_issues" class="anchor"></a>Issues

1\) Is the pipeline struct name too long?

**RESOLVED**: It fits with the naming convention.

2\) Separate W scaling section or fold into coordinate transformations?

**RESOLVED**: Leaving it as its own section for now.

## <a href="#_examples" class="anchor"></a>Examples

``` c
VkViewport viewports[4];
VkRect2D scissors[4];
VkViewportWScalingNV scalings[4];

for (int i = 0; i < 4; i++) {
    int x = (i & 2) ? 0 : currentWindowWidth / 2;
    int y = (i & 1) ? 0 : currentWindowHeight / 2;

    viewports[i].x = 0;
    viewports[i].y = 0;
    viewports[i].width = currentWindowWidth;
    viewports[i].height = currentWindowHeight;
    viewports[i].minDepth = 0.0f;
    viewports[i].maxDepth = 1.0f;

    scissors[i].offset.x = x;
    scissors[i].offset.y = y;
    scissors[i].extent.width = currentWindowWidth/2;
    scissors[i].extent.height = currentWindowHeight/2;

    const float factor = 0.15;
    scalings[i].xcoeff = ((i & 2) ? -1.0 : 1.0) * factor;
    scalings[i].ycoeff = ((i & 1) ? -1.0 : 1.0) * factor;
}

VkPipelineViewportWScalingStateCreateInfoNV vpWScalingStateInfo = { VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV };

vpWScalingStateInfo.viewportWScalingEnable = VK_TRUE;
vpWScalingStateInfo.viewportCount = 4;
vpWScalingStateInfo.pViewportWScalings = &scalings[0];

VkPipelineViewportStateCreateInfo vpStateInfo = { VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO };
vpStateInfo.viewportCount = 4;
vpStateInfo.pViewports = &viewports[0];
vpStateInfo.scissorCount = 4;
vpStateInfo.pScissors = &scissors[0];
vpStateInfo.pNext = &vpWScalingStateInfo;
```

Example shader to read from a w-scaled texture:

``` c
// Vertex Shader
// Draw a triangle that covers the whole screen
const vec4 positions[3] = vec4[3](vec4(-1, -1, 0, 1),
                                  vec4( 3, -1, 0, 1),
                                  vec4(-1,  3, 0, 1));
out vec2 uv;
void main()
{
    vec4 pos = positions[ gl_VertexID ];
    gl_Position = pos;
    uv = pos.xy;
}

// Fragment Shader
uniform sampler2D tex;
uniform float xcoeff;
uniform float ycoeff;
out vec4 Color;
in vec2 uv;

void main()
{
    // Handle uv as if upper right quadrant
    vec2 uvabs = abs(uv);

    // unscale: transform w-scaled image into an unscaled image
    //   scale: transform unscaled image int a w-scaled image
    float unscale = 1.0 / (1 + xcoeff * uvabs.x + xcoeff * uvabs.y);
    //float scale = 1.0 / (1 - xcoeff * uvabs.x - xcoeff * uvabs.y);

    vec2 P = vec2(unscale * uvabs.x, unscale * uvabs.y);

    // Go back to the right quadrant
    P *= sign(uv);

    Color = texture(tex, P * 0.5 + 0.5);
}
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-02-15 (Eric Werness)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_clip_space_w_scaling"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
