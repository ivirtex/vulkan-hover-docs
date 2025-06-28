# VkDisplayModeStereoPropertiesNV(3) Manual Page

## Name

VkDisplayModeStereoPropertiesNV - Structure describing the stereo properties of a display mode



## [](#_c_specification)C Specification

The `VkDisplayModeStereoPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_display_stereo
typedef struct VkDisplayModeStereoPropertiesNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           hdmi3DSupported;
} VkDisplayModeStereoPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `hdmi3DSupported` indicates whether this display mode can be used for a display surface configured for `VK_DISPLAY_SURFACE_STEREO_TYPE_HDMI_3D_NV`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplayModeStereoPropertiesNV-sType-sType)VUID-VkDisplayModeStereoPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_MODE_STEREO_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_display\_stereo](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_display_stereo.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayModeStereoPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0