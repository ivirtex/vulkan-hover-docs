# VkCommandBufferResetFlagBits(3) Manual Page

## Name

VkCommandBufferResetFlagBits - Bitmask controlling behavior of a command
buffer reset



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[vkResetCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetCommandBuffer.html)::`flags`, controlling
the reset operation, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkCommandBufferResetFlagBits {
    VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT = 0x00000001,
} VkCommandBufferResetFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COMMAND_BUFFER_RESET_RELEASE_RESOURCES_BIT` specifies that most or
  all memory resources currently owned by the command buffer **should**
  be returned to the parent command pool. If this flag is not set, then
  the command buffer **may** hold onto memory resources and reuse them
  when recording commands. `commandBuffer` is moved to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">initial state</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBufferResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferResetFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferResetFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
