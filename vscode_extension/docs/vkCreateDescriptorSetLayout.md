# vkCreateDescriptorSetLayout(3) Manual Page

## Name

vkCreateDescriptorSetLayout - Create a new descriptor set layout



## [](#_c_specification)C Specification

To create descriptor set layout objects, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateDescriptorSetLayout(
    VkDevice                                    device,
    const VkDescriptorSetLayoutCreateInfo*      pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorSetLayout*                      pSetLayout);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the descriptor set layout.
- `pCreateInfo` is a pointer to a [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html) structure specifying the state of the descriptor set layout object.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pSetLayout` is a pointer to a [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handle in which the resulting descriptor set layout object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateDescriptorSetLayout-support-09582)VUID-vkCreateDescriptorSetLayout-support-09582  
  If the descriptor layout exceeds the limits reported through the [physical device limits](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits), then [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSupport.html) **must** have returned [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupport.html) with `support` equal to `VK_TRUE` for `pCreateInfo`

Valid Usage (Implicit)

- [](#VUID-vkCreateDescriptorSetLayout-device-parameter)VUID-vkCreateDescriptorSetLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateDescriptorSetLayout-pCreateInfo-parameter)VUID-vkCreateDescriptorSetLayout-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html) structure
- [](#VUID-vkCreateDescriptorSetLayout-pAllocator-parameter)VUID-vkCreateDescriptorSetLayout-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDescriptorSetLayout-pSetLayout-parameter)VUID-vkCreateDescriptorSetLayout-pSetLayout-parameter  
  `pSetLayout` **must** be a valid pointer to a [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDescriptorSetLayout).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0