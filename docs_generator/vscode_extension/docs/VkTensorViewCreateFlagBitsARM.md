# VkTensorViewCreateFlagBitsARM(3) Manual Page

## Name

VkTensorViewCreateFlagBitsARM - Bitmask specifying additional parameters of an tensor view



## [](#_c_specification)C Specification

Bits which **can** be set in [VkTensorViewCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateInfoARM.html)::`flags`, specifying additional parameters of an tensor, are:

```c++
// Provided by VK_ARM_tensors
// Flag bits for VkTensorViewCreateFlagBitsARM
typedef VkFlags64 VkTensorViewCreateFlagBitsARM;
// Provided by VK_EXT_descriptor_buffer with VK_ARM_tensors
static const VkTensorViewCreateFlagBitsARM VK_TENSOR_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM = 0x00000001ULL;
```

## [](#_description)Description

- `VK_TENSOR_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM` specifies that the tensor view **can** be used with descriptor buffers when capturing and replaying (e.g. for trace capture and replay), see [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) for more detail.

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkTensorViewCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateFlagsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorViewCreateFlagBitsARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0