# VkQueueFlagBits(3) Manual Page

## Name

VkQueueFlagBits - Bitmask specifying capabilities of queues in a queue
family



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)::`queueFlags`,
indicating capabilities of queues in a queue family are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkQueueFlagBits {
    VK_QUEUE_GRAPHICS_BIT = 0x00000001,
    VK_QUEUE_COMPUTE_BIT = 0x00000002,
    VK_QUEUE_TRANSFER_BIT = 0x00000004,
    VK_QUEUE_SPARSE_BINDING_BIT = 0x00000008,
  // Provided by VK_VERSION_1_1
    VK_QUEUE_PROTECTED_BIT = 0x00000010,
  // Provided by VK_KHR_video_decode_queue
    VK_QUEUE_VIDEO_DECODE_BIT_KHR = 0x00000020,
  // Provided by VK_KHR_video_encode_queue
    VK_QUEUE_VIDEO_ENCODE_BIT_KHR = 0x00000040,
  // Provided by VK_NV_optical_flow
    VK_QUEUE_OPTICAL_FLOW_BIT_NV = 0x00000100,
} VkQueueFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_QUEUE_GRAPHICS_BIT` specifies that queues in this queue family
  support graphics operations.

- `VK_QUEUE_COMPUTE_BIT` specifies that queues in this queue family
  support compute operations.

- `VK_QUEUE_TRANSFER_BIT` specifies that queues in this queue family
  support transfer operations.

- `VK_QUEUE_SPARSE_BINDING_BIT` specifies that queues in this queue
  family support sparse memory management operations (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sparsememory"
  target="_blank" rel="noopener">Sparse Resources</a>). If any of the
  sparse resource features are enabled, then at least one queue family
  **must** support this bit.

- `VK_QUEUE_VIDEO_DECODE_BIT_KHR` specifies that queues in this queue
  family support <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-decode-operations"
  target="_blank" rel="noopener">video decode operations</a>.

- `VK_QUEUE_VIDEO_ENCODE_BIT_KHR` specifies that queues in this queue
  family support <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-encode-operations"
  target="_blank" rel="noopener">video encode operations</a>.

- `VK_QUEUE_OPTICAL_FLOW_BIT_NV` specifies that queues in this queue
  family support optical flow operations.

- `VK_QUEUE_PROTECTED_BIT` specifies that queues in this queue family
  support the `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT` bit. (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-protected-memory"
  target="_blank" rel="noopener">Protected Memory</a>). If the physical
  device supports the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
  target="_blank" rel="noopener"><code>protectedMemory</code></a>
  feature, at least one of its queue families **must** support this bit.

If an implementation exposes any queue family that supports graphics
operations, at least one queue family of at least one physical device
exposed by the implementation **must** support both graphics and compute
operations.

Furthermore, if the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
target="_blank" rel="noopener"><code>protectedMemory</code></a> physical
device feature is supported, then at least one queue family of at least
one physical device exposed by the implementation **must** support
graphics operations, compute operations, and protected memory
operations.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>All commands that are allowed on a queue that supports transfer
operations are also allowed on a queue that supports either graphics or
compute operations. Thus, if the capabilities of a queue family include
<code>VK_QUEUE_GRAPHICS_BIT</code> or <code>VK_QUEUE_COMPUTE_BIT</code>,
then reporting the <code>VK_QUEUE_TRANSFER_BIT</code> capability
separately for that queue family is <strong>optional</strong>.</p></td>
</tr>
</tbody>
</table>

For further details see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-queues"
target="_blank" rel="noopener">Queues</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkQueueFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
