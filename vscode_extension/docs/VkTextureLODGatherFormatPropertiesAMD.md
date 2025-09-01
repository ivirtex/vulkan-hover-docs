# VkTextureLODGatherFormatPropertiesAMD(3) Manual Page

## Name

VkTextureLODGatherFormatPropertiesAMD - Structure informing whether or not texture gather bias/LOD functionality is supported for a given image format and a given physical device.



## [](#_c_specification)C Specification

To determine if texture gather functions that take explicit LOD and/or bias argument values **can** be used with a given image format, add a [VkTextureLODGatherFormatPropertiesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTextureLODGatherFormatPropertiesAMD.html) structure to the `pNext` chain of the [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure in a call to `vkGetPhysicalDeviceImageFormatProperties2`.

The `VkTextureLODGatherFormatPropertiesAMD` structure is defined as:

```c++
// Provided by VK_AMD_texture_gather_bias_lod
typedef struct VkTextureLODGatherFormatPropertiesAMD {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           supportsTextureGatherLODBiasAMD;
} VkTextureLODGatherFormatPropertiesAMD;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `supportsTextureGatherLODBiasAMD` tells if the image format can be used with texture gather bias/LOD functions, as introduced by the `VK_AMD_texture_gather_bias_lod` extension. This field is set by the implementation. An application-specified value is ignored.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkTextureLODGatherFormatPropertiesAMD-sType-sType)VUID-VkTextureLODGatherFormatPropertiesAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD`

## [](#_see_also)See Also

[VK\_AMD\_texture\_gather\_bias\_lod](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_texture_gather_bias_lod.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTextureLODGatherFormatPropertiesAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0