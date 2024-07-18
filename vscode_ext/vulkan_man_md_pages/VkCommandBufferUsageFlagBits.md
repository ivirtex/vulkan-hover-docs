# VkCommandBufferUsageFlagBits(3) Manual Page

## Name

VkCommandBufferUsageFlagBits - Bitmask specifying usage behavior for
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`flags`,
specifying usage behavior for a command buffer, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkCommandBufferUsageFlagBits {
    VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT = 0x00000001,
    VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT = 0x00000002,
    VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT = 0x00000004,
} VkCommandBufferUsageFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT` specifies that each
  recording of the command buffer will only be submitted once, and the
  command buffer will be reset and recorded again between each
  submission.

- `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` specifies that a
  secondary command buffer is considered to be entirely inside a render
  pass. If this is a primary command buffer, then this bit is ignored.

- `VK_COMMAND_BUFFER_USAGE_SIMULTANEOUS_USE_BIT` specifies that a
  command buffer **can** be resubmitted to any queue of the same queue
  family while it is in the *pending state*, and recorded into multiple
  primary command buffers.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferUsageFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
