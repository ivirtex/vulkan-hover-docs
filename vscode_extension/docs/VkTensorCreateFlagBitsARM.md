# VkTensorCreateFlagBitsARM(3) Manual Page

## Name

VkTensorCreateFlagBitsARM - Bitmask specifying additional parameters of a tensor



## [](#_c_specification)C Specification

Bits which **can** be set in [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`flags`, specifying additional parameters of a tensor, are:

```c++
// Provided by VK_ARM_tensors
// Flag bits for VkTensorCreateFlagBitsARM
typedef VkFlags64 VkTensorCreateFlagBitsARM;
static const VkTensorCreateFlagBitsARM VK_TENSOR_CREATE_MUTABLE_FORMAT_BIT_ARM = 0x00000001ULL;
static const VkTensorCreateFlagBitsARM VK_TENSOR_CREATE_PROTECTED_BIT_ARM = 0x00000002ULL;
// Provided by VK_EXT_descriptor_buffer with VK_ARM_tensors
static const VkTensorCreateFlagBitsARM VK_TENSOR_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM = 0x00000004ULL;
```

## [](#_description)Description

- `VK_TENSOR_CREATE_MUTABLE_FORMAT_BIT_ARM` specifies that the tensor **can** be used to create a `VkTensorViewARM` with a different format from the tensor.
- `VK_TENSOR_CREATE_PROTECTED_BIT_ARM` specifies that the tensor is a protected tensor.
- `VK_TENSOR_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM` specifies that the tensor **can** be used with descriptor buffers when capturing and replaying (e.g. for trace capture and replay), see [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) for more detail.

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkTensorCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorCreateFlagBitsARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0