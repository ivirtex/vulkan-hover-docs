# VkBufferCreateFlagBits(3) Manual Page

## Name

VkBufferCreateFlagBits - Bitmask specifying additional parameters of a
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`flags`, specifying
additional parameters of a buffer, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkBufferCreateFlagBits {
    VK_BUFFER_CREATE_SPARSE_BINDING_BIT = 0x00000001,
    VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT = 0x00000002,
    VK_BUFFER_CREATE_SPARSE_ALIASED_BIT = 0x00000004,
  // Provided by VK_VERSION_1_1
    VK_BUFFER_CREATE_PROTECTED_BIT = 0x00000008,
  // Provided by VK_VERSION_1_2
    VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT = 0x00000010,
  // Provided by VK_EXT_descriptor_buffer
    VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT = 0x00000020,
  // Provided by VK_KHR_video_maintenance1
    VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR = 0x00000040,
  // Provided by VK_EXT_buffer_device_address
    VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT = VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT,
  // Provided by VK_KHR_buffer_device_address
    VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR = VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT,
} VkBufferCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` specifies that the buffer will
  be backed using sparse memory binding.

- `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` specifies that the buffer
  **can** be partially backed using sparse memory binding. Buffers
  created with this flag **must** also be created with the
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` flag.

- `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT` specifies that the buffer will
  be backed using sparse memory binding with memory ranges that might
  also simultaneously be backing another buffer (or another portion of
  the same buffer). Buffers created with this flag **must** also be
  created with the `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` flag.

- `VK_BUFFER_CREATE_PROTECTED_BIT` specifies that the buffer is a
  protected buffer.

- `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` specifies that
  the buffer’s address **can** be saved and reused on a subsequent run
  (e.g. for trace capture and replay), see
  [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)
  for more detail.

- `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` specifies
  that the buffer **can** be used with descriptor buffers when capturing
  and replaying (e.g. for trace capture and replay), see
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  for more detail.

- `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR` specifies that
  the buffer **can** be used in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding"
  target="_blank" rel="noopener">video coding operations</a> without
  having to specify at buffer creation time the set of video profiles
  the buffer will be used with.

See <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sparsememory-sparseresourcefeatures"
target="_blank" rel="noopener">Sparse Resource Features</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features"
target="_blank" rel="noopener">Physical Device Features</a> for details
of the sparse memory features supported on a device.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
