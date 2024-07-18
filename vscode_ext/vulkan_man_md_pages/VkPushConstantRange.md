# VkPushConstantRange(3) Manual Page

## Name

VkPushConstantRange - Structure specifying a push constant range



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPushConstantRange` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPushConstantRange {
    VkShaderStageFlags    stageFlags;
    uint32_t              offset;
    uint32_t              size;
} VkPushConstantRange;
```

## <a href="#_members" class="anchor"></a>Members

- `stageFlags` is a set of stage flags describing the shader stages that
  will access a range of push constants. If a particular stage is not
  included in the range, then accessing members of that range of push
  constants from the corresponding shader stage will return undefined
  values.

- `offset` and `size` are the start offset and size, respectively,
  consumed by the range. Both `offset` and `size` are in units of bytes
  and **must** be a multiple of 4. The layout of the push constant
  variables is specified in the shader.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPushConstantRange-offset-00294"
  id="VUID-VkPushConstantRange-offset-00294"></a>
  VUID-VkPushConstantRange-offset-00294  
  `offset` **must** be less than
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize`

- <a href="#VUID-VkPushConstantRange-offset-00295"
  id="VUID-VkPushConstantRange-offset-00295"></a>
  VUID-VkPushConstantRange-offset-00295  
  `offset` **must** be a multiple of `4`

- <a href="#VUID-VkPushConstantRange-size-00296"
  id="VUID-VkPushConstantRange-size-00296"></a>
  VUID-VkPushConstantRange-size-00296  
  `size` **must** be greater than `0`

- <a href="#VUID-VkPushConstantRange-size-00297"
  id="VUID-VkPushConstantRange-size-00297"></a>
  VUID-VkPushConstantRange-size-00297  
  `size` **must** be a multiple of `4`

- <a href="#VUID-VkPushConstantRange-size-00298"
  id="VUID-VkPushConstantRange-size-00298"></a>
  VUID-VkPushConstantRange-size-00298  
  `size` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize` minus `offset`

Valid Usage (Implicit)

- <a href="#VUID-VkPushConstantRange-stageFlags-parameter"
  id="VUID-VkPushConstantRange-stageFlags-parameter"></a>
  VUID-VkPushConstantRange-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a href="#VUID-VkPushConstantRange-stageFlags-requiredbitmask"
  id="VUID-VkPushConstantRange-stageFlags-requiredbitmask"></a>
  VUID-VkPushConstantRange-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html),
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPushConstantRange"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
