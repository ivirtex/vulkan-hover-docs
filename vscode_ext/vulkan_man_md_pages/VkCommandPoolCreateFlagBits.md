# VkCommandPoolCreateFlagBits(3) Manual Page

## Name

VkCommandPoolCreateFlagBits - Bitmask specifying usage behavior for a
command pool



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkCommandPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateInfo.html)::`flags`,
specifying usage behavior for a command pool, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkCommandPoolCreateFlagBits {
    VK_COMMAND_POOL_CREATE_TRANSIENT_BIT = 0x00000001,
    VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT = 0x00000002,
  // Provided by VK_VERSION_1_1
    VK_COMMAND_POOL_CREATE_PROTECTED_BIT = 0x00000004,
} VkCommandPoolCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COMMAND_POOL_CREATE_TRANSIENT_BIT` specifies that command buffers
  allocated from the pool will be short-lived, meaning that they will be
  reset or freed in a relatively short timeframe. This flag **may** be
  used by the implementation to control memory allocation behavior
  within the pool.

- `VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT` allows any command
  buffer allocated from a pool to be individually reset to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">initial state</a>; either by calling
  [vkResetCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetCommandBuffer.html), or via the implicit
  reset when calling [vkBeginCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBeginCommandBuffer.html).
  If this flag is not set on a pool, then `vkResetCommandBuffer`
  **must** not be called for any command buffer allocated from that
  pool.

- `VK_COMMAND_POOL_CREATE_PROTECTED_BIT` specifies that command buffers
  allocated from the pool are protected command buffers.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandPoolCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandPoolCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
