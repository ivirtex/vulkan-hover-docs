# VkCoarseSampleLocationNV(3) Manual Page

## Name

VkCoarseSampleLocationNV - Structure specifying parameters controlling shading rate image usage



## [](#_c_specification)C Specification

The `VkCoarseSampleLocationNV` structure identifies a specific pixel and [sample index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask) for one of the coverage samples in a fragment that is larger than one pixel. This structure is defined as:

```c++
// Provided by VK_NV_shading_rate_image
typedef struct VkCoarseSampleLocationNV {
    uint32_t    pixelX;
    uint32_t    pixelY;
    uint32_t    sample;
} VkCoarseSampleLocationNV;
```

## [](#_members)Members

- `pixelX` is added to the x coordinate of the upper-leftmost pixel of each fragment to identify the pixel containing the coverage sample.
- `pixelY` is added to the y coordinate of the upper-leftmost pixel of each fragment to identify the pixel containing the coverage sample.
- `sample` is the number of the coverage sample in the pixel identified by `pixelX` and `pixelY`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCoarseSampleLocationNV-pixelX-02078)VUID-VkCoarseSampleLocationNV-pixelX-02078  
  `pixelX` **must** be less than the width (in pixels) of the fragment
- [](#VUID-VkCoarseSampleLocationNV-pixelY-02079)VUID-VkCoarseSampleLocationNV-pixelY-02079  
  `pixelY` **must** be less than the height (in pixels) of the fragment
- [](#VUID-VkCoarseSampleLocationNV-sample-02080)VUID-VkCoarseSampleLocationNV-sample-02080  
  `sample` **must** be less than the number of coverage samples in each pixel belonging to the fragment

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderCustomNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCoarseSampleLocationNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0