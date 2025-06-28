# VK\_EXT\_shader\_tile\_image(3) Manual Page

## Name

VK\_EXT\_shader\_tile\_image - device extension



## [](#_registered_extension_number)Registered Extension Number

396

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.3](#versions-1.3)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_shader\_tile\_image](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_shader_tile_image.html)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharaldfredriksen-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_tile_image%5D%20%40janharaldfredriksen-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_tile_image%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_shader\_tile\_image](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_shader_tile_image.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-03-23

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_shader_tile_image`](https://raw.githubusercontent.com/KhronosGroup/GLSL/main/extensions/ext/GLSL_EXT_shader_tile_image.txt)

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

## [](#_description)Description

This extension allows fragment shader invocations to read color, depth and stencil values at their pixel location in rasterization order. The functionality is only available when using dynamic render passes introduced by VK\_KHR\_dynamic\_rendering. Example use cases are programmable blending and deferred shading.

See [fragment shader tile image reads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-shader-tileimage-reads) for more information.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderTileImageFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderTileImageFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderTileImagePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderTileImagePropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_TILE_IMAGE_EXTENSION_NAME`
- `VK_EXT_SHADER_TILE_IMAGE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TILE_IMAGE_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_TILE_IMAGE_PROPERTIES_EXT`

## [](#_issues)Issues

None.

## [](#_examples)Examples

Color read example.

```c
layout( location = 0 /* aliased to color attachment 0 */ ) tileImageEXT highp attachmentEXT color0;
layout( location = 1 /* aliased to color attachment 1 */ ) tileImageEXT highp attachmentEXT color1;

layout( location = 0 ) out vec4 fragColor;

void main()
{
    vec4 value = colorAttachmentReadEXT(color0) + colorAttachmentReadEXT(color1);
    fragColor = value;
}
```

Depth &amp; Stencil read example.

```c
void main()
{
    // read sample 0: works for non-MSAA or MSAA targets
    highp float last_depth = depthAttachmentReadEXT();
    lowp uint last_stencil = stencilAttachmentReadEXT();

    //..
}
```

## [](#_version_history)Version History

- Revision 1, 2023-03-23 (Sandeep Kakarlapudi)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_tile_image)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0