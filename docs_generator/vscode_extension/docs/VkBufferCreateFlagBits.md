# VkBufferCreateFlagBits(3) Manual Page

## Name

VkBufferCreateFlagBits - Bitmask specifying additional parameters of a buffer



## [](#_c_specification)C Specification

Bits which **can** be set in [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`flags`, specifying additional parameters of a buffer, are:

```c++
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

## [](#_description)Description

- `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` specifies that the buffer will be backed using sparse memory binding.
- `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` specifies that the buffer **can** be partially backed using sparse memory binding. Buffers created with this flag **must** also be created with the `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` flag.
- `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT` specifies that the buffer will be backed using sparse memory binding with memory ranges that might also simultaneously be backing another buffer (or another portion of the same buffer). Buffers created with this flag **must** also be created with the `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` flag.
- `VK_BUFFER_CREATE_PROTECTED_BIT` specifies that the buffer is a protected buffer.
- `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` specifies that the buffer’s address **can** be saved and reused on a subsequent run (e.g. for trace capture and replay), see [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html) for more detail.
- `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` specifies that the buffer **can** be used with descriptor buffers when capturing and replaying (e.g. for trace capture and replay), see [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) for more detail.
- `VK_BUFFER_CREATE_VIDEO_PROFILE_INDEPENDENT_BIT_KHR` specifies that the buffer **can** be used in [video coding operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-coding) without having to specify at buffer creation time the set of video profiles the buffer will be used with.

See [Sparse Resource Features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#sparsememory-sparseresourcefeatures) and [Physical Device Features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features) for details of the sparse memory features supported on a device.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferCreateFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0