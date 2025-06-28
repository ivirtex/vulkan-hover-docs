# VkPipelineRasterizationStateStreamCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationStateStreamCreateInfoEXT - Structure defining the geometry stream used for rasterization



## [](#_c_specification)C Specification

The vertex stream used for rasterization is specified by adding a `VkPipelineRasterizationStateStreamCreateInfoEXT` structure to the `pNext` chain of a [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html) structure.

The `VkPipelineRasterizationStateStreamCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_transform_feedback
typedef struct VkPipelineRasterizationStateStreamCreateInfoEXT {
    VkStructureType                                     sType;
    const void*                                         pNext;
    VkPipelineRasterizationStateStreamCreateFlagsEXT    flags;
    uint32_t                                            rasterizationStream;
} VkPipelineRasterizationStateStreamCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `rasterizationStream` is the vertex stream selected for rasterization.

## [](#_description)Description

If this structure is not present, `rasterizationStream` is assumed to be zero.

Valid Usage

- [](#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-geometryStreams-02324)VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-geometryStreams-02324  
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`geometryStreams` **must** be enabled
- [](#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02325)VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02325  
  `rasterizationStream` **must** be less than [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackStreams`
- [](#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02326)VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-rasterizationStream-02326  
  `rasterizationStream` **must** be zero if `VkPhysicalDeviceTransformFeedbackPropertiesEXT`::`transformFeedbackRasterizationStreamSelect` is `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-sType-sType)VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_STATE_STREAM_CREATE_INFO_EXT`
- [](#VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-flags-zerobitmask)VUID-VkPipelineRasterizationStateStreamCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_EXT\_transform\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_transform_feedback.html), [VkPipelineRasterizationStateStreamCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateStreamCreateFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRasterizationStateStreamCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0