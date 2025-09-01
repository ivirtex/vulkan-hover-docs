# VkTensorARM(3) Manual Page

## Name

VkTensorARM - Opaque handle to a tensor object



## [](#_c_specification)C Specification

Tensors represent multidimensional arrays of data. Tensors **can** be used by binding them to pipelines via descriptor sets, or by directly specifying them as parameters to certain commands.

Tensors are represented by `VkTensorARM` handles:

```c++
// Provided by VK_ARM_tensors
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkTensorARM)
```

## [](#_description)Description

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VkBindTensorMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindTensorMemoryInfoARM.html), [VkCopyTensorInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyTensorInfoARM.html), [VkFrameBoundaryTensorsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryTensorsARM.html), [VkMemoryDedicatedAllocateInfoTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfoTensorARM.html), [VkTensorCaptureDescriptorDataInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCaptureDescriptorDataInfoARM.html), [VkTensorMemoryBarrierARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryBarrierARM.html), [VkTensorMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorMemoryRequirementsInfoARM.html), [VkTensorViewCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateInfoARM.html), [vkCreateTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateTensorARM.html), [vkDestroyTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyTensorARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0