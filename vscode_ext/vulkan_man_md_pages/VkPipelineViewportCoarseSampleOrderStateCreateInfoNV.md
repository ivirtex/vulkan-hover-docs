# VkPipelineViewportCoarseSampleOrderStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportCoarseSampleOrderStateCreateInfoNV - Structure
specifying parameters controlling sample order in coarse fragments



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
includes a `VkPipelineViewportCoarseSampleOrderStateCreateInfoNV`
structure, then that structure includes parameters controlling the order
of coverage samples in fragments larger than one pixel.

The `VkPipelineViewportCoarseSampleOrderStateCreateInfoNV` structure is
defined as:

``` c
// Provided by VK_NV_shading_rate_image
typedef struct VkPipelineViewportCoarseSampleOrderStateCreateInfoNV {
    VkStructureType                       sType;
    const void*                           pNext;
    VkCoarseSampleOrderTypeNV             sampleOrderType;
    uint32_t                              customSampleOrderCount;
    const VkCoarseSampleOrderCustomNV*    pCustomSampleOrders;
} VkPipelineViewportCoarseSampleOrderStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `sampleOrderType` specifies the mechanism used to order coverage
  samples in fragments larger than one pixel.

- `customSampleOrderCount` specifies the number of custom sample
  orderings to use when ordering coverage samples.

- `pCustomSampleOrders` is a pointer to an array of
  `customSampleOrderCount`
  [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html)
  structures, each structure specifying the coverage sample order for a
  single combination of fragment area and coverage sample count.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, `sampleOrderType` is considered to be
`VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV`.

If `sampleOrderType` is `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`, the
coverage sample order used for any combination of fragment area and
coverage sample count not enumerated in `pCustomSampleOrders` will be
identical to that used for `VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV`.

If the pipeline was created with
`VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV`, the contents of this
structure (if present) are ignored, and the coverage sample order is
instead specified by
[vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoarseSampleOrderNV.html).

Valid Usage

- <a
  href="#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-02072"
  id="VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-02072"></a>
  VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-02072  
  If `sampleOrderType` is not `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`,
  `customSamplerOrderCount` **must** be `0`

- <a
  href="#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-02234"
  id="VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-02234"></a>
  VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-02234  
  The array `pCustomSampleOrders` **must** not contain two structures
  with matching values for both the `shadingRate` and `sampleCount`
  members

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV`

- <a
  href="#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-parameter"
  id="VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-parameter"></a>
  VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-parameter  
  `sampleOrderType` **must** be a valid
  [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderTypeNV.html) value

- <a
  href="#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-parameter"
  id="VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-parameter"></a>
  VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-parameter  
  If `customSampleOrderCount` is not `0`, `pCustomSampleOrders` **must**
  be a valid pointer to an array of `customSampleOrderCount` valid
  [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html),
[VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderTypeNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineViewportCoarseSampleOrderStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
