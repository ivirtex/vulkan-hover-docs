# vkCreateDescriptorSetLayout(3) Manual Page

## Name

vkCreateDescriptorSetLayout - Create a new descriptor set layout



## <a href="#_c_specification" class="anchor"></a>C Specification

To create descriptor set layout objects, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateDescriptorSetLayout(
    VkDevice                                    device,
    const VkDescriptorSetLayoutCreateInfo*      pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDescriptorSetLayout*                      pSetLayout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the descriptor set layout.

- `pCreateInfo` is a pointer to a
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
  structure specifying the state of the descriptor set layout object.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pSetLayout` is a pointer to a
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handle in which
  the resulting descriptor set layout object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateDescriptorSetLayout-support-09582"
  id="VUID-vkCreateDescriptorSetLayout-support-09582"></a>
  VUID-vkCreateDescriptorSetLayout-support-09582  
  If the descriptor layout exceeds the limits reported through the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits"
  target="_blank" rel="noopener">physical device limits</a>, then
  [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html)
  **must** have returned
  [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html) with
  `support` equal to `VK_TRUE` for `pCreateInfo`

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDescriptorSetLayout-device-parameter"
  id="VUID-vkCreateDescriptorSetLayout-device-parameter"></a>
  VUID-vkCreateDescriptorSetLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateDescriptorSetLayout-pCreateInfo-parameter"
  id="VUID-vkCreateDescriptorSetLayout-pCreateInfo-parameter"></a>
  VUID-vkCreateDescriptorSetLayout-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
  structure

- <a href="#VUID-vkCreateDescriptorSetLayout-pAllocator-parameter"
  id="VUID-vkCreateDescriptorSetLayout-pAllocator-parameter"></a>
  VUID-vkCreateDescriptorSetLayout-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateDescriptorSetLayout-pSetLayout-parameter"
  id="VUID-vkCreateDescriptorSetLayout-pSetLayout-parameter"></a>
  VUID-vkCreateDescriptorSetLayout-pSetLayout-parameter  
  `pSetLayout` **must** be a valid pointer to a
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDescriptorSetLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
