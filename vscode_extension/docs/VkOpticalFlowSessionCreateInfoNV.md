# VkOpticalFlowSessionCreateInfoNV(3) Manual Page

## Name

VkOpticalFlowSessionCreateInfoNV - Structure specifying parameters of a newly created optical flow session



## [](#_c_specification)C Specification

The [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowSessionCreateInfoNV {
    VkStructureType                      sType;
    void*                                pNext;
    uint32_t                             width;
    uint32_t                             height;
    VkFormat                             imageFormat;
    VkFormat                             flowVectorFormat;
    VkFormat                             costFormat;
    VkOpticalFlowGridSizeFlagsNV         outputGridSize;
    VkOpticalFlowGridSizeFlagsNV         hintGridSize;
    VkOpticalFlowPerformanceLevelNV      performanceLevel;
    VkOpticalFlowSessionCreateFlagsNV    flags;
} VkOpticalFlowSessionCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `width` is the width in pixels of the input or reference frame to be bound to this optical flow session.
- `height` is the height in pixels of the input or reference frame to be bound to this optical flow session.
- `imageFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the input and reference frame to be bound to this optical flow session.
- `flowVectorFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the flow vector maps (output or hint) to be bound to this optical flow session.
- `costFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the cost maps to be bound to this optical flow session.
- `outputGridSize` is exactly one bit of [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html) specifying the grid size of the output flow and cost maps to be bound to this optical flow session. The size of the output flow and cost maps is determined by `VkOpticalFlowSessionCreateInfoNV`::`width` and `VkOpticalFlowSessionCreateInfoNV`::`height` divided by `VkOpticalFlowSessionCreateInfoNV`::`outputGridSize`.
- `hintGridSize` is one exactly bit of [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html) specifying the grid size of the hint flow vector maps to be bound to this optical flow session. The size of the hint maps is determined by `VkOpticalFlowSessionCreateInfoNV`::`width` and `VkOpticalFlowSessionCreateInfoNV`::`height` divided by `VkOpticalFlowSessionCreateInfoNV`::`hintGridSize`.
- `performanceLevel` is the [VkOpticalFlowPerformanceLevelNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowPerformanceLevelNV.html) used for this optical flow session.
- `flags` are the [VkOpticalFlowSessionCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateFlagsNV.html) used for this optical flow session.

## [](#_description)Description

Valid Usage

- [](#VUID-VkOpticalFlowSessionCreateInfoNV-width-07581)VUID-VkOpticalFlowSessionCreateInfoNV-width-07581  
  `width` **must** be greater than or equal to `VkPhysicalDeviceOpticalFlowPropertiesNV`::`minWidth` and less than or equal to `VkPhysicalDeviceOpticalFlowPropertiesNV`::`maxWidth`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-height-07582)VUID-VkOpticalFlowSessionCreateInfoNV-height-07582  
  `height` **must** be greater than or equal to `VkPhysicalDeviceOpticalFlowPropertiesNV`::`minHeight` and less than or equal to `VkPhysicalDeviceOpticalFlowPropertiesNV`::`maxHeight`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-07583)VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-07583  
  `imageFormat` **must** be one of the formats returned by [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html) for `VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-07584)VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-07584  
  `flowVectorFormat` **must** be one of the formats returned by [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html) for `VK_OPTICAL_FLOW_USAGE_OUTPUT_BIT_NV`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-07585)VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-07585  
  `costFormat` **must** be one of the formats returned by [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html) for `VK_OPTICAL_FLOW_USAGE_COST_BIT_NV` if `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_COST_BIT_NV` is set in `flags`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-07586)VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-07586  
  `outputGridSize` **must** be exactly one of the bits reported in `VkPhysicalDeviceOpticalFlowPropertiesNV`::`supportedOutputGridSizes`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-07587)VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-07587  
  `hintGridSize` **must** be exactly one of the bits reported in `VkPhysicalDeviceOpticalFlowPropertiesNV`::`supportedHintGridSizes` if `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_HINT_BIT_NV` is set in `flags`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07588)VUID-VkOpticalFlowSessionCreateInfoNV-flags-07588  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_HINT_BIT_NV` **must** not be set in `flags` if `VkPhysicalDeviceOpticalFlowPropertiesNV`::`hintSupported` is `VK_FALSE`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07589)VUID-VkOpticalFlowSessionCreateInfoNV-flags-07589  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_COST_BIT_NV` **must** not be set in `flags` if `VkPhysicalDeviceOpticalFlowPropertiesNV`::`costSupported` is `VK_FALSE`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07590)VUID-VkOpticalFlowSessionCreateInfoNV-flags-07590  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_GLOBAL_FLOW_BIT_NV` **must** not be set in `flags` if `VkPhysicalDeviceOpticalFlowPropertiesNV`::`globalFlowSupported` is `VK_FALSE`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07591)VUID-VkOpticalFlowSessionCreateInfoNV-flags-07591  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ALLOW_REGIONS_BIT_NV` **must** not be set in `flags` if `VkPhysicalDeviceOpticalFlowPropertiesNV`::`maxNumRegionsOfInterest` is 0
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07592)VUID-VkOpticalFlowSessionCreateInfoNV-flags-07592  
  `VK_OPTICAL_FLOW_SESSION_CREATE_BOTH_DIRECTIONS_BIT_NV` **must** not be set in `flags` if `VkPhysicalDeviceOpticalFlowPropertiesNV`::`bidirectionalFlowSupported` is `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkOpticalFlowSessionCreateInfoNV-sType-sType)VUID-VkOpticalFlowSessionCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_OPTICAL_FLOW_SESSION_CREATE_INFO_NV`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-pNext-pNext)VUID-VkOpticalFlowSessionCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkOpticalFlowSessionCreatePrivateDataInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreatePrivateDataInfoNV.html)
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-sType-unique)VUID-VkOpticalFlowSessionCreateInfoNV-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-parameter)VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-parameter  
  `imageFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-parameter)VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-parameter  
  `flowVectorFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-parameter)VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-parameter  
  If `costFormat` is not `0`, `costFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-parameter)VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-parameter  
  `outputGridSize` **must** be a valid combination of [VkOpticalFlowGridSizeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagBitsNV.html) values
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-requiredbitmask)VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-requiredbitmask  
  `outputGridSize` **must** not be `0`
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-parameter)VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-parameter  
  `hintGridSize` **must** be a valid combination of [VkOpticalFlowGridSizeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagBitsNV.html) values
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-performanceLevel-parameter)VUID-VkOpticalFlowSessionCreateInfoNV-performanceLevel-parameter  
  If `performanceLevel` is not `0`, `performanceLevel` **must** be a valid [VkOpticalFlowPerformanceLevelNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowPerformanceLevelNV.html) value
- [](#VUID-VkOpticalFlowSessionCreateInfoNV-flags-parameter)VUID-VkOpticalFlowSessionCreateInfoNV-flags-parameter  
  `flags` **must** be a valid combination of [VkOpticalFlowSessionCreateFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html), [VkOpticalFlowPerformanceLevelNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowPerformanceLevelNV.html), [VkOpticalFlowSessionCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateOpticalFlowSessionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkOpticalFlowSessionCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0