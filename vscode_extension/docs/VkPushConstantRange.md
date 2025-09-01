# VkPushConstantRange(3) Manual Page

## Name

VkPushConstantRange - Structure specifying a push constant range



## [](#_c_specification)C Specification

The `VkPushConstantRange` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPushConstantRange {
    VkShaderStageFlags    stageFlags;
    uint32_t              offset;
    uint32_t              size;
} VkPushConstantRange;
```

## [](#_members)Members

- `stageFlags` is a set of stage flags describing the shader stages that will access a range of push constants. If a particular stage is not included in the range, then accessing members of that range of push constants from the corresponding shader stage will return undefined values.
- `offset` and `size` are the start offset and size, respectively, consumed by the range. Both `offset` and `size` are in units of bytes and **must** be a multiple of 4. The layout of the push constant variables is specified in the shader.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPushConstantRange-offset-00294)VUID-VkPushConstantRange-offset-00294  
  `offset` **must** be less than `VkPhysicalDeviceLimits`::`maxPushConstantsSize`
- [](#VUID-VkPushConstantRange-offset-00295)VUID-VkPushConstantRange-offset-00295  
  `offset` **must** be a multiple of `4`
- [](#VUID-VkPushConstantRange-size-00296)VUID-VkPushConstantRange-size-00296  
  `size` **must** be greater than `0`
- [](#VUID-VkPushConstantRange-size-00297)VUID-VkPushConstantRange-size-00297  
  `size` **must** be a multiple of `4`
- [](#VUID-VkPushConstantRange-size-00298)VUID-VkPushConstantRange-size-00298  
  `size` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxPushConstantsSize` minus `offset`

Valid Usage (Implicit)

- [](#VUID-VkPushConstantRange-stageFlags-parameter)VUID-VkPushConstantRange-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkPushConstantRange-stageFlags-requiredbitmask)VUID-VkPushConstantRange-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkIndirectCommandsPushConstantTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsPushConstantTokenEXT.html), [VkIndirectExecutionSetShaderInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderInfoEXT.html), [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html), [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPushConstantRange).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0