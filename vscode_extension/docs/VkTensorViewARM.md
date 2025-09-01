# VkTensorViewARM(3) Manual Page

## Name

VkTensorViewARM - Opaque handle to an tensor view object



## [](#_c_specification)C Specification

Tensor objects are not directly accessed by pipelines for reading or writing tensor data. Instead, *tensor views* representing the tensor subresources and containing additional metadata are used for that purpose. Views **must** be created on tensors of compatible types.

Tensor views are represented by `VkTensorViewARM` handles:

```c++
// Provided by VK_ARM_tensors
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkTensorViewARM)
```

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VkDescriptorGetTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetTensorInfoARM.html), [VkTensorViewCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCaptureDescriptorDataInfoARM.html), [VkWriteDescriptorSetTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetTensorARM.html), [vkCreateTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateTensorViewARM.html), [vkDestroyTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyTensorViewARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorViewARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0