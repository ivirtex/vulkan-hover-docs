# VkPipelineRasterizationStateStreamCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationStateStreamCreateInfoEXT - Structure defining the
geometry stream used for rasterization



## <a href="#_c_specification" class="anchor"></a>C Specification

The vertex stream used for rasterization is specified by adding a
`VkPipelineRasterizationStateStreamCreateInfoEXT` structure to the
`pNext` chain of a
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
structure.

The `VkPipelineRasterizationStateStreamCreateInfoEXT` structure is
defined as:

``` c
// Provided by VK_EXT_transform_feedback
typedef struct VkPipelineRasterizationStateStreamCreateInfoEXT {
    VkStructureType                                     sType;
    const void*                                         pNext;
    VkPipelineRasterizationStateStreamCreateFlagsEXT    flags;
    uint32_t                                            rasterizationStream;
} VkPipelineRasterizationStateStreamCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `rasterizationStream` is the vertex stream selected for rasterization.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, `rasterizationStream` is assumed to be
zero.

Valid Usage

- <a
  href="#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-geometryStreams-02324"
  id="VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-geometryStreams-02324"></a>
  VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-geometryStreams-02324  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`geometryStreams`
  **must** be enabled

- <a
  href="#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02325"
  id="VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02325"></a>
  VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02325  
  `rasterizationStream` **must** be less than
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackStreams`

- <a
  href="#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02326"
  id="VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02326"></a>
  VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02326  
  `rasterizationStream` **must** be zero if
  `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`transformFeedbackRasterizationStreamSelect`
  is `VK_FALSE`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_transform_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_transform_feedback.html),
[VkPipelineRasterizationStateStreamCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateStreamCreateFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRasterizationStateStreamCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
