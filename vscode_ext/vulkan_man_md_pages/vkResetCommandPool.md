# vkResetCommandPool(3) Manual Page

## Name

vkResetCommandPool - Reset a command pool



## <a href="#_c_specification" class="anchor"></a>C Specification

To reset a command pool, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkResetCommandPool(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    VkCommandPoolResetFlags                     flags);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the command pool.

- `commandPool` is the command pool to reset.

- `flags` is a bitmask of
  [VkCommandPoolResetFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolResetFlagBits.html)
  controlling the reset operation.

## <a href="#_description" class="anchor"></a>Description

Resetting a command pool recycles all of the resources from all of the
command buffers allocated from the command pool back to the command
pool. All command buffers that have been allocated from the command pool
are put in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">initial state</a>.

Any primary command buffer allocated from another
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html) that is in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">recording or executable state</a> and has
a secondary command buffer allocated from `commandPool` recorded into
it, becomes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid</a>.

Valid Usage

- <a href="#VUID-vkResetCommandPool-commandPool-00040"
  id="VUID-vkResetCommandPool-commandPool-00040"></a>
  VUID-vkResetCommandPool-commandPool-00040  
  All `VkCommandBuffer` objects allocated from `commandPool` **must**
  not be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

Valid Usage (Implicit)

- <a href="#VUID-vkResetCommandPool-device-parameter"
  id="VUID-vkResetCommandPool-device-parameter"></a>
  VUID-vkResetCommandPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkResetCommandPool-commandPool-parameter"
  id="VUID-vkResetCommandPool-commandPool-parameter"></a>
  VUID-vkResetCommandPool-commandPool-parameter  
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html)
  handle

- <a href="#VUID-vkResetCommandPool-flags-parameter"
  id="VUID-vkResetCommandPool-flags-parameter"></a>
  VUID-vkResetCommandPool-flags-parameter  
  `flags` **must** be a valid combination of
  [VkCommandPoolResetFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolResetFlagBits.html) values

- <a href="#VUID-vkResetCommandPool-commandPool-parent"
  id="VUID-vkResetCommandPool-commandPool-parent"></a>
  VUID-vkResetCommandPool-commandPool-parent  
  `commandPool` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html),
[VkCommandPoolResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolResetFlags.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkResetCommandPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
