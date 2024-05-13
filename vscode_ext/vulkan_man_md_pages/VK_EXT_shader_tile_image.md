# VK_EXT_shader_tile_image(3) Manual Page

## Name

VK_EXT_shader_tile_image - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

396

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.3](#versions-1.3)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_shader_tile_image](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_shader_tile_image.html)

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_tile_image%5D%20@janharaldfredriksen-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_tile_image%20extension*"
  target="_blank"
  rel="nofollow noopener"><em></em>janharaldfredriksen-arm</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_shader_tile_image](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_tile_image.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-03-23

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_EXT_shader_tile_image`](https://raw.githubusercontent.com/KhronosGroup/GLSL/main/extensions/ext/GLSL_EXT_shader_tile_image.txt)

**Contributors**  
- Sandeep Kakarlapudi, Arm

- Jan-Harald Fredriksen, Arm

- James Fitzpatrick, Imagination

- Andrew Garrard, Imagination

- Jeff Leger, Qualcomm

- Huilong Wang, Huawei

- Graeme Leese, Broadcom

- Hans-Kristian Arntzen, Valve

- Tobias Hector, AMD

- Jeff Bolz, NVIDIA

- Shahbaz Youssefi, Google

## <a href="#_description" class="anchor"></a>Description

This extension allows fragment shader invocations to read color, depth
and stencil values at their pixel location in rasterization order. The
functionality is only available when using dynamic render passes
introduced by VK_KHR_dynamic_rendering. Example use cases are
programmable blending and deferred shading.

See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-shader-tileimage-reads"
target="_blank" rel="noopener">fragment shader tile image reads</a> for
more information.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderTileImageFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderTileImageFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderTileImagePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderTileImagePropertiesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SHADER_TILE_IMAGE_EXTENSION_NAME`

- `VK_EXT_SHADER_TILE_IMAGE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TILE_IMAGE_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TILE_IMAGE_PROPERTIES_EXT`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_examples" class="anchor"></a>Examples

Color read example.

``` c
layout( location = 0 /* aliased to color attachment 0 */ ) tileImageEXT highp attachmentEXT color0;
layout( location = 1 /* aliased to color attachment 1 */ ) tileImageEXT highp attachmentEXT color1;

layout( location = 0 ) out vec4 fragColor;

void main()
{
    vec4 value = colorAttachmentReadEXT(color0) + colorAttachmentReadEXT(color1);
    fragColor = value;
}
```

Depth & Stencil read example.

``` c
void main()
{
    // read sample 0: works for non-MSAA or MSAA targets
    highp float last_depth = depthAttachmentReadEXT();
    lowp uint last_stencil = stencilAttachmentReadEXT();

    //..
}
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-03-23 (Sandeep Kakarlapudi)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_shader_tile_image"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
