# VkTensorViewCaptureDescriptorDataInfoARM(3) Manual Page

## Name

VkTensorViewCaptureDescriptorDataInfoARM - Structure specifying a tensor view for descriptor capture



## [](#_c_specification)C Specification

Information about the tensor view to get descriptor buffer capture data for is passed in a `VkTensorViewCaptureDescriptorDataInfoARM` structure:

```c++
// Provided by VK_EXT_descriptor_buffer with VK_ARM_tensors
typedef struct VkTensorViewCaptureDescriptorDataInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    VkTensorViewARM    tensorView;
} VkTensorViewCaptureDescriptorDataInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `tensorView` is the `VkTensorViewARM` handle of the tensor view to get opaque capture data for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkTensorViewCaptureDescriptorDataInfoARM-tensorView-09709)VUID-VkTensorViewCaptureDescriptorDataInfoARM-tensorView-09709  
  If `tensorView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) then `tensorView` **must** have been created with `VK_TENSOR_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM` set in [VkTensorViewCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateInfoARM.html)::`flags`

Valid Usage (Implicit)

- [](#VUID-VkTensorViewCaptureDescriptorDataInfoARM-sType-sType)VUID-VkTensorViewCaptureDescriptorDataInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_VIEW_CAPTURE_DESCRIPTOR_DATA_INFO_ARM`
- [](#VUID-VkTensorViewCaptureDescriptorDataInfoARM-pNext-pNext)VUID-VkTensorViewCaptureDescriptorDataInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkTensorViewCaptureDescriptorDataInfoARM-tensorView-parameter)VUID-VkTensorViewCaptureDescriptorDataInfoARM-tensorView-parameter  
  `tensorView` **must** be a valid [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) handle

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html), [vkGetTensorViewOpaqueCaptureDescriptorDataARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetTensorViewOpaqueCaptureDescriptorDataARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorViewCaptureDescriptorDataInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0