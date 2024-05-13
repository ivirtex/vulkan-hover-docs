# VkRenderingFlagBits(3) Manual Page

## Name

VkRenderingFlagBits - Bitmask specifying additional properties of a
dynamic render pass instance



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`flags` describing additional
properties of the render pass are:

``` c
// Provided by VK_VERSION_1_3
typedef enum VkRenderingFlagBits {
    VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT = 0x00000001,
    VK_RENDERING_SUSPENDING_BIT = 0x00000002,
    VK_RENDERING_RESUMING_BIT = 0x00000004,
  // Provided by VK_EXT_nested_command_buffer
    VK_RENDERING_CONTENTS_INLINE_BIT_EXT = 0x00000010,
  // Provided by VK_EXT_legacy_dithering with (VK_KHR_dynamic_rendering or VK_VERSION_1_3) and VK_KHR_maintenance5
    VK_RENDERING_ENABLE_LEGACY_DITHERING_BIT_EXT = 0x00000008,
    VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT_KHR = VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT,
    VK_RENDERING_SUSPENDING_BIT_KHR = VK_RENDERING_SUSPENDING_BIT,
    VK_RENDERING_RESUMING_BIT_KHR = VK_RENDERING_RESUMING_BIT,
} VkRenderingFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_dynamic_rendering
typedef VkRenderingFlagBits VkRenderingFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT` specifies that
  draw calls for the render pass instance will be recorded in secondary
  command buffers. If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
  target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
  feature is enabled, the draw calls **can** come from both inline and
  [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html).

- `VK_RENDERING_RESUMING_BIT` specifies that the render pass instance is
  resuming an earlier suspended render pass instance.

- `VK_RENDERING_SUSPENDING_BIT` specifies that the render pass instance
  will be suspended.

- `VK_RENDERING_ENABLE_LEGACY_DITHERING_BIT_EXT` specifies that <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-legacy-dithering"
  target="_blank" rel="noopener">Legacy Dithering</a> is enabled for the
  render pass instance.

- `VK_RENDERING_CONTENTS_INLINE_BIT_EXT` specifies that draw calls for
  the render pass instance **can** be recorded inline within the current
  command buffer. When the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
  target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
  feature is enabled this **can** be combined with the
  `VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT` bit to allow
  draw calls to be recorded both inline and in secondary command
  buffers.

The contents of `pRenderingInfo` **must** match between suspended render
pass instances and the render pass instances that resume them, other
than the presence or absence of the `VK_RENDERING_RESUMING_BIT`,
`VK_RENDERING_SUSPENDING_BIT`, and
`VK_RENDERING_CONTENTS_SECONDARY_COMMAND_BUFFERS_BIT` flags. No action
or synchronization commands, or other render pass instances, are allowed
between suspending and resuming render pass instances.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkRenderingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderingFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
