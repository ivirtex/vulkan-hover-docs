# VK\_AMD\_texture\_gather\_bias\_lod(3) Manual Page

## Name

VK\_AMD\_texture\_gather\_bias\_lod - device extension



## [](#_registered_extension_number)Registered Extension Number

42

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_AMD\_texture\_gather\_bias\_lod](https://github.khronos.org/SPIRV-Registry/extensions/AMD/SPV_AMD_texture_gather_bias_lod.html)

## [](#_contact)Contact

- Rex Xu [\[GitHub\]amdrexu](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_texture_gather_bias_lod%5D%20%40amdrexu%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_texture_gather_bias_lod%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-03-21

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_AMD_texture_gather_bias_lod`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_texture_gather_bias_lod.txt)

**Contributors**

- Dominik Witczak, AMD
- Daniel Rakos, AMD
- Graham Sellers, AMD
- Matthaeus G. Chajdas, AMD
- Qun Lin, AMD
- Rex Xu, AMD
- Timothy Lottes, AMD

## [](#_description)Description

This extension adds two related features.

Firstly, support for the following SPIR-V extension in Vulkan is added:

- `SPV_AMD_texture_gather_bias_lod`

Secondly, the extension allows the application to query which formats can be used together with the new function prototypes introduced by the SPIR-V extension.

## [](#_new_structures)New Structures

- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html):
  
  - [VkTextureLODGatherFormatPropertiesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTextureLODGatherFormatPropertiesAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_TEXTURE_GATHER_BIAS_LOD_EXTENSION_NAME`
- `VK_AMD_TEXTURE_GATHER_BIAS_LOD_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`ImageGatherBiasLodAMD`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ImageGatherBiasLodAMD)

## [](#_examples)Examples

```c++
struct VkTextureLODGatherFormatPropertiesAMD
{
    VkStructureType sType;
    const void*     pNext;
    VkBool32        supportsTextureGatherLODBiasAMD;
};

// ----------------------------------------------------------------------------------------
// How to detect if an image format can be used with the new function prototypes.
VkPhysicalDeviceImageFormatInfo2   formatInfo;
VkImageFormatProperties2           formatProps;
VkTextureLODGatherFormatPropertiesAMD textureLODGatherSupport;

textureLODGatherSupport.sType = VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD;
textureLODGatherSupport.pNext = nullptr;

formatInfo.sType  = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2;
formatInfo.pNext  = nullptr;
formatInfo.format = ...;
formatInfo.type   = ...;
formatInfo.tiling = ...;
formatInfo.usage  = ...;
formatInfo.flags  = ...;

formatProps.sType = VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2;
formatProps.pNext = &textureLODGatherSupport;

vkGetPhysicalDeviceImageFormatProperties2(physical_device, &formatInfo, &formatProps);

if (textureLODGatherSupport.supportsTextureGatherLODBiasAMD == VK_TRUE)
{
    // physical device supports SPV_AMD_texture_gather_bias_lod for the specified
    // format configuration.
}
else
{
    // physical device does not support SPV_AMD_texture_gather_bias_lod for the
    // specified format configuration.
}
```

## [](#_version_history)Version History

- Revision 1, 2017-03-21 (Dominik Witczak)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_texture_gather_bias_lod)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0