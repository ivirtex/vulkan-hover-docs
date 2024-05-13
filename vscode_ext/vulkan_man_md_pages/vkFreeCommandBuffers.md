# vkFreeCommandBuffers(3) Manual Page

## Name

vkFreeCommandBuffers - Free command buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

To free command buffers, call:

``` c
// Provided by VK_VERSION_1_0
void vkFreeCommandBuffers(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    uint32_t                                    commandBufferCount,
    const VkCommandBuffer*                      pCommandBuffers);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the command pool.

- `commandPool` is the command pool from which the command buffers were
  allocated.

- `commandBufferCount` is the length of the `pCommandBuffers` array.

- `pCommandBuffers` is a pointer to an array of handles of command
  buffers to free.

## <a href="#_description" class="anchor"></a>Description

Any primary command buffer that is in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">recording or executable state</a> and has
any element of `pCommandBuffers` recorded into it, becomes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">invalid</a>.

Valid Usage

- <a href="#VUID-vkFreeCommandBuffers-pCommandBuffers-00047"
  id="VUID-vkFreeCommandBuffers-pCommandBuffers-00047"></a>
  VUID-vkFreeCommandBuffers-pCommandBuffers-00047  
  All elements of `pCommandBuffers` **must** not be in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

- <a href="#VUID-vkFreeCommandBuffers-pCommandBuffers-00048"
  id="VUID-vkFreeCommandBuffers-pCommandBuffers-00048"></a>
  VUID-vkFreeCommandBuffers-pCommandBuffers-00048  
  `pCommandBuffers` **must** be a valid pointer to an array of
  `commandBufferCount` `VkCommandBuffer` handles, each element of which
  **must** either be a valid handle or `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkFreeCommandBuffers-device-parameter"
  id="VUID-vkFreeCommandBuffers-device-parameter"></a>
  VUID-vkFreeCommandBuffers-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkFreeCommandBuffers-commandPool-parameter"
  id="VUID-vkFreeCommandBuffers-commandPool-parameter"></a>
  VUID-vkFreeCommandBuffers-commandPool-parameter  
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html)
  handle

- <a href="#VUID-vkFreeCommandBuffers-commandBufferCount-arraylength"
  id="VUID-vkFreeCommandBuffers-commandBufferCount-arraylength"></a>
  VUID-vkFreeCommandBuffers-commandBufferCount-arraylength  
  `commandBufferCount` **must** be greater than `0`

- <a href="#VUID-vkFreeCommandBuffers-commandPool-parent"
  id="VUID-vkFreeCommandBuffers-commandPool-parent"></a>
  VUID-vkFreeCommandBuffers-commandPool-parent  
  `commandPool` **must** have been created, allocated, or retrieved from
  `device`

- <a href="#VUID-vkFreeCommandBuffers-pCommandBuffers-parent"
  id="VUID-vkFreeCommandBuffers-pCommandBuffers-parent"></a>
  VUID-vkFreeCommandBuffers-pCommandBuffers-parent  
  Each element of `pCommandBuffers` that is a valid handle **must** have
  been created, allocated, or retrieved from `commandPool`

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized

- Host access to each member of `pCommandBuffers` **must** be externally
  synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkFreeCommandBuffers"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
