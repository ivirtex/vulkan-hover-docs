# VkSubpassContents(3) Manual Page

## Name

VkSubpassContents - Specify how commands in the first subpass of a
render pass are provided



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html)::`contents`,
specifying how the commands in the first subpass will be provided, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkSubpassContents {
    VK_SUBPASS_CONTENTS_INLINE = 0,
    VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS = 1,
  // Provided by VK_EXT_nested_command_buffer
    VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_EXT = 1000451000,
} VkSubpassContents;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SUBPASS_CONTENTS_INLINE` specifies that the contents of the
  subpass will be recorded inline in the primary command buffer, and
  secondary command buffers **must** not be executed within the subpass.

- `VK_SUBPASS_CONTENTS_SECONDARY_COMMAND_BUFFERS` specifies that the
  contents are recorded in secondary command buffers that will be called
  from the primary command buffer, and
  [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html) is the only valid
  command in the command buffer until
  [vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html) or
  [vkCmdEndRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEndRenderPass.html).

- `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_EXT`
  specifies that the contents of the subpass **can** be recorded both
  inline and in secondary command buffers executed from this command
  buffer with [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSubpassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassBeginInfo.html),
[vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html),
[vkCmdNextSubpass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdNextSubpass.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassContents"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
