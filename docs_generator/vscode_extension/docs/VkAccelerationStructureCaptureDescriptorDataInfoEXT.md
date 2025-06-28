# VkAccelerationStructureCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkAccelerationStructureCaptureDescriptorDataInfoEXT - Structure specifying an acceleration structure for descriptor capture



## [](#_c_specification)C Specification

Information about the acceleration structure to get descriptor buffer capture data for is passed in a `VkAccelerationStructureCaptureDescriptorDataInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer with VK_KHR_acceleration_structure or VK_NV_ray_tracing
typedef struct VkAccelerationStructureCaptureDescriptorDataInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    VkAccelerationStructureKHR    accelerationStructure;
    VkAccelerationStructureNV     accelerationStructureNV;
} VkAccelerationStructureCaptureDescriptorDataInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `accelerationStructure` is the `VkAccelerationStructureKHR` handle of the acceleration structure to get opaque capture data for.
- `accelerationStructureNV` is the `VkAccelerationStructureNV` handle of the acceleration structure to get opaque capture data for.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08091)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08091  
  If `accelerationStructure` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) then `accelerationStructure` **must** have been created with `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`createFlags`
- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08092)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08092  
  If `accelerationStructureNV` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) then `accelerationStructureNV` **must** have been created with `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoNV.html)::`info.flags`
- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08093)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-08093  
  If `accelerationStructure` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) then `accelerationStructureNV` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08094)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-08094  
  If `accelerationStructureNV` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) then `accelerationStructure` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-sType-sType)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-pNext-pNext)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-parameter)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructure-parameter  
  If `accelerationStructure` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `accelerationStructure` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-parameter)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-accelerationStructureNV-parameter  
  If `accelerationStructureNV` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `accelerationStructureNV` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-commonparent)VUID-VkAccelerationStructureCaptureDescriptorDataInfoEXT-commonparent  
  Both of `accelerationStructure`, and `accelerationStructureNV` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureCaptureDescriptorDataInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0