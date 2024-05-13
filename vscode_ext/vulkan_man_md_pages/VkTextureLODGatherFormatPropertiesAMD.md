# VkTextureLODGatherFormatPropertiesAMD(3) Manual Page

## Name

VkTextureLODGatherFormatPropertiesAMD - Structure informing whether or
not texture gather bias/LOD functionality is supported for a given image
format and a given physical device.



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine if texture gather functions that take explicit LOD and/or
bias argument values **can** be used with a given image format, add a
[VkTextureLODGatherFormatPropertiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTextureLODGatherFormatPropertiesAMD.html)
structure to the `pNext` chain of the
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure in a
call to `vkGetPhysicalDeviceImageFormatProperties2`.

The `VkTextureLODGatherFormatPropertiesAMD` structure is defined as:

``` c
// Provided by VK_AMD_texture_gather_bias_lod
typedef struct VkTextureLODGatherFormatPropertiesAMD {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           supportsTextureGatherLODBiasAMD;
} VkTextureLODGatherFormatPropertiesAMD;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `supportsTextureGatherLODBiasAMD` tells if the image format can be
  used with texture gather bias/LOD functions, as introduced by the
  [`VK_AMD_texture_gather_bias_lod`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_texture_gather_bias_lod.html)
  extension. This field is set by the implementation. User-specified
  value is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkTextureLODGatherFormatPropertiesAMD-sType-sType"
  id="VUID-VkTextureLODGatherFormatPropertiesAMD-sType-sType"></a>
  VUID-VkTextureLODGatherFormatPropertiesAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_TEXTURE_LOD_GATHER_FORMAT_PROPERTIES_AMD`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_texture_gather_bias_lod](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_texture_gather_bias_lod.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkTextureLODGatherFormatPropertiesAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
