# VkTensorCaptureDescriptorDataInfoARM(3) Manual Page

## Name

VkTensorCaptureDescriptorDataInfoARM - Structure specifying a tensor for descriptor capture



## [](#_c_specification)C Specification

Information about the tensor to get descriptor buffer capture data for is passed in a `VkTensorCaptureDescriptorDataInfoARM` structure:

```c++
// Provided by VK_EXT_descriptor_buffer with VK_ARM_tensors
typedef struct VkTensorCaptureDescriptorDataInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkTensorARM        tensor;
} VkTensorCaptureDescriptorDataInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensor` is the `VkTensorARM` handle of the tensor to get opaque capture data for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkTensorCaptureDescriptorDataInfoARM-tensor-09705)VUID-VkTensorCaptureDescriptorDataInfoARM-tensor-09705  
  If `tensor` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) then `tensor` **must** have been created with `VK_TENSOR_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM` set in [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`flags`

Valid Usage (Implicit)

- [](#VUID-VkTensorCaptureDescriptorDataInfoARM-sType-sType)VUID-VkTensorCaptureDescriptorDataInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_CAPTURE_DESCRIPTOR_DATA_INFO_ARM`
- [](#VUID-VkTensorCaptureDescriptorDataInfoARM-pNext-pNext)VUID-VkTensorCaptureDescriptorDataInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkTensorCaptureDescriptorDataInfoARM-tensor-parameter)VUID-VkTensorCaptureDescriptorDataInfoARM-tensor-parameter  
  `tensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [vkGetTensorOpaqueCaptureDescriptorDataARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorOpaqueCaptureDescriptorDataARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorCaptureDescriptorDataInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0