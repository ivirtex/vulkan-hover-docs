# vkResetDescriptorPool(3) Manual Page

## Name

vkResetDescriptorPool - Resets a descriptor pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

To return all descriptor sets allocated from a given pool to the pool,
rather than freeing individual descriptor sets, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkResetDescriptorPool(
    VkDevice                                    device,
    VkDescriptorPool                            descriptorPool,
    VkDescriptorPoolResetFlags                  flags);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the descriptor pool.

- `descriptorPool` is the descriptor pool to be reset.

- `flags` is reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Resetting a descriptor pool recycles all of the resources from all of
the descriptor sets allocated from the descriptor pool back to the
descriptor pool, and the descriptor sets are implicitly freed.

Valid Usage

- <a href="#VUID-vkResetDescriptorPool-descriptorPool-00313"
  id="VUID-vkResetDescriptorPool-descriptorPool-00313"></a>
  VUID-vkResetDescriptorPool-descriptorPool-00313  
  All uses of `descriptorPool` (via any allocated descriptor sets)
  **must** have completed execution

Valid Usage (Implicit)

- <a href="#VUID-vkResetDescriptorPool-device-parameter"
  id="VUID-vkResetDescriptorPool-device-parameter"></a>
  VUID-vkResetDescriptorPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkResetDescriptorPool-descriptorPool-parameter"
  id="VUID-vkResetDescriptorPool-descriptorPool-parameter"></a>
  VUID-vkResetDescriptorPool-descriptorPool-parameter  
  `descriptorPool` **must** be a valid
  [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html) handle

- <a href="#VUID-vkResetDescriptorPool-flags-zerobitmask"
  id="VUID-vkResetDescriptorPool-flags-zerobitmask"></a>
  VUID-vkResetDescriptorPool-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-vkResetDescriptorPool-descriptorPool-parent"
  id="VUID-vkResetDescriptorPool-descriptorPool-parent"></a>
  VUID-vkResetDescriptorPool-descriptorPool-parent  
  `descriptorPool` **must** have been created, allocated, or retrieved
  from `device`

Host Synchronization

- Host access to `descriptorPool` **must** be externally synchronized

- Host access to any `VkDescriptorSet` objects allocated from
  `descriptorPool` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPool.html),
[VkDescriptorPoolResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolResetFlags.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkResetDescriptorPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
