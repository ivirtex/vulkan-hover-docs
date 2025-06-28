# vkGetDescriptorSetLayoutSupport(3) Manual Page

## Name

vkGetDescriptorSetLayoutSupport - Query whether a descriptor set layout can be created



## [](#_c_specification)C Specification

To query information about whether a descriptor set layout **can** be created, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetDescriptorSetLayoutSupport(
    VkDevice                                    device,
    const VkDescriptorSetLayoutCreateInfo*      pCreateInfo,
    VkDescriptorSetLayoutSupport*               pSupport);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance3
void vkGetDescriptorSetLayoutSupportKHR(
    VkDevice                                    device,
    const VkDescriptorSetLayoutCreateInfo*      pCreateInfo,
    VkDescriptorSetLayoutSupport*               pSupport);
```

## [](#_parameters)Parameters

- `device` is the logical device that would create the descriptor set layout.
- `pCreateInfo` is a pointer to a [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html) structure specifying the state of the descriptor set layout object.
- `pSupport` is a pointer to a [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupport.html) structure, in which information about support for the descriptor set layout object is returned.

## [](#_description)Description

Some implementations have limitations on what fits in a descriptor set which are not easily expressible in terms of existing limits like `maxDescriptorSet`\*, for example if all descriptor types share a limited space in memory but each descriptor is a different size or alignment. This command returns information about whether a descriptor set satisfies this limit. If the descriptor set layout satisfies the [VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance3Properties.html)::`maxPerSetDescriptors` limit, this command is guaranteed to return `VK_TRUE` in [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupport.html)::`supported`. If the descriptor set layout exceeds the [VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance3Properties.html)::`maxPerSetDescriptors` limit, whether the descriptor set layout is supported is implementation-dependent and **may** depend on whether the descriptor sizes and alignments cause the layout to exceed an internal limit.

This command does not consider other limits such as `maxPerStageDescriptor`\*, and so a descriptor set layout that is supported according to this command **must** still satisfy the pipeline layout limits such as `maxPerStageDescriptor`* in order to be used in a pipeline layout.

Note

This is a `VkDevice` query rather than `VkPhysicalDevice` because the answer **may** depend on enabled features.

Valid Usage (Implicit)

- [](#VUID-vkGetDescriptorSetLayoutSupport-device-parameter)VUID-vkGetDescriptorSetLayoutSupport-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDescriptorSetLayoutSupport-pCreateInfo-parameter)VUID-vkGetDescriptorSetLayoutSupport-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html) structure
- [](#VUID-vkGetDescriptorSetLayoutSupport-pSupport-parameter)VUID-vkGetDescriptorSetLayoutSupport-pSupport-parameter  
  `pSupport` **must** be a valid pointer to a [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupport.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html), [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupport.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDescriptorSetLayoutSupport)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0