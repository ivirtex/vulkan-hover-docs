# VkCoarseSampleLocationNV(3) Manual Page

## Name

VkCoarseSampleLocationNV - Structure specifying parameters controlling
shading rate image usage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCoarseSampleLocationNV` structure identifies a specific pixel and
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
target="_blank" rel="noopener">sample index</a> for one of the coverage
samples in a fragment that is larger than one pixel. This structure is
defined as:

``` c
// Provided by VK_NV_shading_rate_image
typedef struct VkCoarseSampleLocationNV {
    uint32_t    pixelX;
    uint32_t    pixelY;
    uint32_t    sample;
} VkCoarseSampleLocationNV;
```

## <a href="#_members" class="anchor"></a>Members

- `pixelX` is added to the x coordinate of the upper-leftmost pixel of
  each fragment to identify the pixel containing the coverage sample.

- `pixelY` is added to the y coordinate of the upper-leftmost pixel of
  each fragment to identify the pixel containing the coverage sample.

- `sample` is the number of the coverage sample in the pixel identified
  by `pixelX` and `pixelY`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCoarseSampleLocationNV-pixelX-02078"
  id="VUID-VkCoarseSampleLocationNV-pixelX-02078"></a>
  VUID-VkCoarseSampleLocationNV-pixelX-02078  
  `pixelX` **must** be less than the width (in pixels) of the fragment

- <a href="#VUID-VkCoarseSampleLocationNV-pixelY-02079"
  id="VUID-VkCoarseSampleLocationNV-pixelY-02079"></a>
  VUID-VkCoarseSampleLocationNV-pixelY-02079  
  `pixelY` **must** be less than the height (in pixels) of the fragment

- <a href="#VUID-VkCoarseSampleLocationNV-sample-02080"
  id="VUID-VkCoarseSampleLocationNV-sample-02080"></a>
  VUID-VkCoarseSampleLocationNV-sample-02080  
  `sample` **must** be less than the number of coverage samples in each
  pixel belonging to the fragment

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCoarseSampleLocationNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
