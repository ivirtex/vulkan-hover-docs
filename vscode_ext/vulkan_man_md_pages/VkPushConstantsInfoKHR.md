# VkPushConstantsInfoKHR(3) Manual Page

## Name

VkPushConstantsInfoKHR - Structure specifying a push constant update
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPushConstantsInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance6
typedef struct VkPushConstantsInfoKHR {
    VkStructureType       sType;
    const void*           pNext;
    VkPipelineLayout      layout;
    VkShaderStageFlags    stageFlags;
    uint32_t              offset;
    uint32_t              size;
    const void*           pValues;
} VkPushConstantsInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `layout` is the pipeline layout used to program the push constant
  updates. If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicPipelineLayout"
  target="_blank" rel="noopener"><code>dynamicPipelineLayout</code></a>
  feature is enabled, `layout` **can** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the layout **must** be
  specified by chaining
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure off the `pNext`

- `stageFlags` is a bitmask of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) specifying the
  shader stages that will use the push constants in the updated range.

- `offset` is the start offset of the push constant range to update, in
  units of bytes.

- `size` is the size of the push constant range to update, in units of
  bytes.

- `pValues` is a pointer to an array of `size` bytes containing the new
  push constant values.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPushConstantsInfoKHR-offset-01795"
  id="VUID-VkPushConstantsInfoKHR-offset-01795"></a>
  VUID-VkPushConstantsInfoKHR-offset-01795  
  For each byte in the range specified by `offset` and `size` and for
  each shader stage in `stageFlags`, there **must** be a push constant
  range in `layout` that includes that byte and that stage

- <a href="#VUID-VkPushConstantsInfoKHR-offset-01796"
  id="VUID-VkPushConstantsInfoKHR-offset-01796"></a>
  VUID-VkPushConstantsInfoKHR-offset-01796  
  For each byte in the range specified by `offset` and `size` and for
  each push constant range that overlaps that byte, `stageFlags`
  **must** include all stages in that push constant range’s
  [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantRange.html)::`stageFlags`

- <a href="#VUID-VkPushConstantsInfoKHR-offset-00368"
  id="VUID-VkPushConstantsInfoKHR-offset-00368"></a>
  VUID-VkPushConstantsInfoKHR-offset-00368  
  `offset` **must** be a multiple of `4`

- <a href="#VUID-VkPushConstantsInfoKHR-size-00369"
  id="VUID-VkPushConstantsInfoKHR-size-00369"></a>
  VUID-VkPushConstantsInfoKHR-size-00369  
  `size` **must** be a multiple of `4`

- <a href="#VUID-VkPushConstantsInfoKHR-offset-00370"
  id="VUID-VkPushConstantsInfoKHR-offset-00370"></a>
  VUID-VkPushConstantsInfoKHR-offset-00370  
  `offset` **must** be less than
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize`

- <a href="#VUID-VkPushConstantsInfoKHR-size-00371"
  id="VUID-VkPushConstantsInfoKHR-size-00371"></a>
  VUID-VkPushConstantsInfoKHR-size-00371  
  `size` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPushConstantsSize` minus `offset`

<!-- -->

- <a href="#VUID-VkPushConstantsInfoKHR-None-09495"
  id="VUID-VkPushConstantsInfoKHR-None-09495"></a>
  VUID-VkPushConstantsInfoKHR-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout)
  feature is not enabled, `layout` **must** be a valid
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkPushConstantsInfoKHR-layout-09496"
  id="VUID-VkPushConstantsInfoKHR-layout-09496"></a>
  VUID-VkPushConstantsInfoKHR-layout-09496  
  If `layout` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `pNext`
  chain **must** include a valid
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkPushConstantsInfoKHR-sType-sType"
  id="VUID-VkPushConstantsInfoKHR-sType-sType"></a>
  VUID-VkPushConstantsInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PUSH_CONSTANTS_INFO_KHR`

- <a href="#VUID-VkPushConstantsInfoKHR-pNext-pNext"
  id="VUID-VkPushConstantsInfoKHR-pNext-pNext"></a>
  VUID-VkPushConstantsInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)

- <a href="#VUID-VkPushConstantsInfoKHR-sType-unique"
  id="VUID-VkPushConstantsInfoKHR-sType-unique"></a>
  VUID-VkPushConstantsInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPushConstantsInfoKHR-layout-parameter"
  id="VUID-VkPushConstantsInfoKHR-layout-parameter"></a>
  VUID-VkPushConstantsInfoKHR-layout-parameter  
  If `layout` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `layout`
  **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkPushConstantsInfoKHR-stageFlags-parameter"
  id="VUID-VkPushConstantsInfoKHR-stageFlags-parameter"></a>
  VUID-VkPushConstantsInfoKHR-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a href="#VUID-VkPushConstantsInfoKHR-stageFlags-requiredbitmask"
  id="VUID-VkPushConstantsInfoKHR-stageFlags-requiredbitmask"></a>
  VUID-VkPushConstantsInfoKHR-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`

- <a href="#VUID-VkPushConstantsInfoKHR-pValues-parameter"
  id="VUID-VkPushConstantsInfoKHR-pValues-parameter"></a>
  VUID-VkPushConstantsInfoKHR-pValues-parameter  
  `pValues` **must** be a valid pointer to an array of `size` bytes

- <a href="#VUID-VkPushConstantsInfoKHR-size-arraylength"
  id="VUID-VkPushConstantsInfoKHR-size-arraylength"></a>
  VUID-VkPushConstantsInfoKHR-size-arraylength  
  `size` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdPushConstants2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushConstants2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPushConstantsInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
