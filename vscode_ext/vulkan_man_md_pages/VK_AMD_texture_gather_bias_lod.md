# VK_AMD_texture_gather_bias_lod(3) Manual Page

## Name

VK_AMD_texture_gather_bias_lod - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

42

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_AMD_texture_gather_bias_lod](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/AMD/SPV_AMD_texture_gather_bias_lod.html)

## <a href="#_contact" class="anchor"></a>Contact

- Rex Xu <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_texture_gather_bias_lod%5D%20@amdrexu%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_texture_gather_bias_lod%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>amdrexu</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-03-21

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_AMD_texture_gather_bias_lod`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_texture_gather_bias_lod.txt)

**Contributors**  
- Dominik Witczak, AMD

- Daniel Rakos, AMD

- Graham Sellers, AMD

- Matthaeus G. Chajdas, AMD

- Qun Lin, AMD

- Rex Xu, AMD

- Timothy Lottes, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds two related features.

Firstly, support for the following SPIR-V extension in Vulkan is added:

- `SPV_AMD_texture_gather_bias_lod`

Secondly, the extension allows the application to query which formats
can be used together with the new function prototypes introduced by the
SPIR-V extension.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html):

  - [VkTextureLODGatherFormatPropertiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTextureLODGatherFormatPropertiesAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_TEXTURE_GATHER_BIAS_LOD_EXTENSION_NAME`

- `VK_AMD_TEXTURE_GATHER_BIAS_LOD_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ImageGatherBiasLodAMD"
  target="_blank" rel="noopener"><code>ImageGatherBiasLodAMD</code></a>

## <a href="#_examples" class="anchor"></a>Examples

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-03-21 (Dominik Witczak)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_texture_gather_bias_lod"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
