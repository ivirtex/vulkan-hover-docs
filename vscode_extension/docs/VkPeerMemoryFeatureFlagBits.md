# VkPeerMemoryFeatureFlagBits(3) Manual Page

## Name

VkPeerMemoryFeatureFlagBits - Bitmask specifying supported peer memory
features



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[vkGetDeviceGroupPeerMemoryFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPeerMemoryFeatures.html)::`pPeerMemoryFeatures`,
indicating supported peer memory features, are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkPeerMemoryFeatureFlagBits {
    VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT = 0x00000001,
    VK_PEER_MEMORY_FEATURE_COPY_DST_BIT = 0x00000002,
    VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT = 0x00000004,
    VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT = 0x00000008,
  // Provided by VK_KHR_device_group
    VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR = VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT,
  // Provided by VK_KHR_device_group
    VK_PEER_MEMORY_FEATURE_COPY_DST_BIT_KHR = VK_PEER_MEMORY_FEATURE_COPY_DST_BIT,
  // Provided by VK_KHR_device_group
    VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT_KHR = VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT,
  // Provided by VK_KHR_device_group
    VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT_KHR = VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT,
} VkPeerMemoryFeatureFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group
typedef VkPeerMemoryFeatureFlagBits VkPeerMemoryFeatureFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT` specifies that the memory
  **can** be accessed as the source of any `vkCmdCopy*` command.

- `VK_PEER_MEMORY_FEATURE_COPY_DST_BIT` specifies that the memory
  **can** be accessed as the destination of any `vkCmdCopy*` command.

- `VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT` specifies that the memory
  **can** be read as any memory access type.

- `VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT` specifies that the memory
  **can** be written as any memory access type. Shader atomics are
  considered to be writes.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The peer memory features of a memory heap also apply to any accesses
that <strong>may</strong> be performed during <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
target="_blank" rel="noopener">image layout transitions</a>.</p></td>
</tr>
</tbody>
</table>

`VK_PEER_MEMORY_FEATURE_COPY_DST_BIT` **must** be supported for all host
local heaps and for at least one device-local memory heap.

If a device does not support a peer memory feature, it is still valid to
use a resource that includes both local and peer memory bindings with
the corresponding access type as long as only the local bindings are
actually accessed. For example, an application doing split-frame
rendering would use framebuffer attachments that include both local and
peer memory bindings, but would scissor the rendering to only update
local memory.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPeerMemoryFeatureFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
