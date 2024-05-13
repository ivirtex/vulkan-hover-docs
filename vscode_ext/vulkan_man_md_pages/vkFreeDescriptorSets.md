# vkFreeDescriptorSets(3) Manual Page

## Name

vkFreeDescriptorSets - Free one or more descriptor sets



## <a href="#_c_specification" class="anchor"></a>C Specification

To free allocated descriptor sets, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkFreeDescriptorSets(
    VkDevice                                    device,
    VkDescriptorPool                            descriptorPool,
    uint32_t                                    descriptorSetCount,
    const VkDescriptorSet*                      pDescriptorSets);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the descriptor pool.

- `descriptorPool` is the descriptor pool from which the descriptor sets
  were allocated.

- `descriptorSetCount` is the number of elements in the
  `pDescriptorSets` array.

- `pDescriptorSets` is a pointer to an array of handles to
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) objects.

## <a href="#_description" class="anchor"></a>Description

After calling `vkFreeDescriptorSets`, all descriptor sets in
`pDescriptorSets` are invalid.

Valid Usage

- <a href="#VUID-vkFreeDescriptorSets-pDescriptorSets-00309"
  id="VUID-vkFreeDescriptorSets-pDescriptorSets-00309"></a>
  VUID-vkFreeDescriptorSets-pDescriptorSets-00309  
  All submitted commands that refer to any element of `pDescriptorSets`
  **must** have completed execution

- <a href="#VUID-vkFreeDescriptorSets-pDescriptorSets-00310"
  id="VUID-vkFreeDescriptorSets-pDescriptorSets-00310"></a>
  VUID-vkFreeDescriptorSets-pDescriptorSets-00310  
  `pDescriptorSets` **must** be a valid pointer to an array of
  `descriptorSetCount` `VkDescriptorSet` handles, each element of which
  **must** either be a valid handle or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkFreeDescriptorSets-descriptorPool-00312"
  id="VUID-vkFreeDescriptorSets-descriptorPool-00312"></a>
  VUID-vkFreeDescriptorSets-descriptorPool-00312  
  `descriptorPool` **must** have been created with the
  `VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT` flag

Valid Usage (Implicit)

- <a href="#VUID-vkFreeDescriptorSets-device-parameter"
  id="VUID-vkFreeDescriptorSets-device-parameter"></a>
  VUID-vkFreeDescriptorSets-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkFreeDescriptorSets-descriptorPool-parameter"
  id="VUID-vkFreeDescriptorSets-descriptorPool-parameter"></a>
  VUID-vkFreeDescriptorSets-descriptorPool-parameter  
  `descriptorPool` **must** be a valid
  [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html) handle

- <a href="#VUID-vkFreeDescriptorSets-descriptorSetCount-arraylength"
  id="VUID-vkFreeDescriptorSets-descriptorSetCount-arraylength"></a>
  VUID-vkFreeDescriptorSets-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`

- <a href="#VUID-vkFreeDescriptorSets-descriptorPool-parent"
  id="VUID-vkFreeDescriptorSets-descriptorPool-parent"></a>
  VUID-vkFreeDescriptorSets-descriptorPool-parent  
  `descriptorPool` **must** have been created, allocated, or retrieved
  from `device`

- <a href="#VUID-vkFreeDescriptorSets-pDescriptorSets-parent"
  id="VUID-vkFreeDescriptorSets-pDescriptorSets-parent"></a>
  VUID-vkFreeDescriptorSets-pDescriptorSets-parent  
  Each element of `pDescriptorSets` that is a valid handle **must** have
  been created, allocated, or retrieved from `descriptorPool`

Host Synchronization

- Host access to `descriptorPool` **must** be externally synchronized

- Host access to each member of `pDescriptorSets` **must** be externally
  synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html),
[VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkFreeDescriptorSets"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
