# VkPushConstantsInfo(3) Manual Page

## Name

VkPushConstantsInfo - Structure specifying a push constant update operation



## [](#_c_specification)C Specification

The `VkPushConstantsInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPushConstantsInfo {
    VkStructureType       sType;
    const void*           pNext;
    VkPipelineLayout      layout;
    VkShaderStageFlags    stageFlags;
    uint32_t              offset;
    uint32_t              size;
    const void*           pValues;
} VkPushConstantsInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance6
typedef VkPushConstantsInfo VkPushConstantsInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `layout` is the pipeline layout used to program the push constant updates. If the [`dynamicPipelineLayout`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicPipelineLayout) feature is enabled, `layout` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the layout **must** be specified by chaining [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure off the `pNext`
- `stageFlags` is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying the shader stages that will use the push constants in the updated range.
- `offset` is the start offset of the push constant range to update, in units of bytes.
- `size` is the size of the push constant range to update, in units of bytes.
- `pValues` is a pointer to an array of `size` bytes containing the new push constant values.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPushConstantsInfo-offset-01795)VUID-VkPushConstantsInfo-offset-01795  
  For each byte in the range specified by `offset` and `size` and for each shader stage in `stageFlags`, there **must** be a push constant range in `layout` that includes that byte and that stage
- [](#VUID-VkPushConstantsInfo-offset-01796)VUID-VkPushConstantsInfo-offset-01796  
  For each byte in the range specified by `offset` and `size` and for each push constant range that overlaps that byte, `stageFlags` **must** include all stages in that push constant range’s [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html)::`stageFlags`
- [](#VUID-VkPushConstantsInfo-offset-00368)VUID-VkPushConstantsInfo-offset-00368  
  `offset` **must** be a multiple of `4`
- [](#VUID-VkPushConstantsInfo-size-00369)VUID-VkPushConstantsInfo-size-00369  
  `size` **must** be a multiple of `4`
- [](#VUID-VkPushConstantsInfo-offset-00370)VUID-VkPushConstantsInfo-offset-00370  
  `offset` **must** be less than `VkPhysicalDeviceLimits`::`maxPushConstantsSize`
- [](#VUID-VkPushConstantsInfo-size-00371)VUID-VkPushConstantsInfo-size-00371  
  `size` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxPushConstantsSize` minus `offset`

<!--THE END-->

- [](#VUID-VkPushConstantsInfo-None-09495)VUID-VkPushConstantsInfo-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout) feature is not enabled, `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkPushConstantsInfo-layout-09496)VUID-VkPushConstantsInfo-layout-09496  
  If `layout` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `pNext` chain **must** include a valid [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) structure

Valid Usage (Implicit)

- [](#VUID-VkPushConstantsInfo-sType-sType)VUID-VkPushConstantsInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PUSH_CONSTANTS_INFO`
- [](#VUID-VkPushConstantsInfo-pNext-pNext)VUID-VkPushConstantsInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html)
- [](#VUID-VkPushConstantsInfo-sType-unique)VUID-VkPushConstantsInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPushConstantsInfo-layout-parameter)VUID-VkPushConstantsInfo-layout-parameter  
  If `layout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) handle
- [](#VUID-VkPushConstantsInfo-stageFlags-parameter)VUID-VkPushConstantsInfo-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkPushConstantsInfo-stageFlags-requiredbitmask)VUID-VkPushConstantsInfo-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`
- [](#VUID-VkPushConstantsInfo-pValues-parameter)VUID-VkPushConstantsInfo-pValues-parameter  
  `pValues` **must** be a valid pointer to an array of `size` bytes
- [](#VUID-VkPushConstantsInfo-size-arraylength)VUID-VkPushConstantsInfo-size-arraylength  
  `size` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdPushConstants2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushConstants2.html), [vkCmdPushConstants2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushConstants2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPushConstantsInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0