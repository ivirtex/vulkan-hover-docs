# VkPhysicalDeviceExternalTensorInfoARM(3) Manual Page

## Name

VkPhysicalDeviceExternalTensorInfoARM - Structure specifying tensor creation parameters.



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExternalTensorInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkPhysicalDeviceExternalTensorInfoARM {
    VkStructureType                       sType;
    const void*                           pNext;
    VkTensorCreateFlagsARM                flags;
    const VkTensorDescriptionARM*         pDescription;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkPhysicalDeviceExternalTensorInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkTensorCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagBitsARM.html) describing additional parameters of the tensor, corresponding to [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`flags`.
- `pDescription` is a [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure describing the tensor, corresponding to [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`pDescription`.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the external memory handle type for which capabilities will be returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExternalTensorInfoARM-sType-sType)VUID-VkPhysicalDeviceExternalTensorInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_TENSOR_INFO_ARM`
- [](#VUID-VkPhysicalDeviceExternalTensorInfoARM-pNext-pNext)VUID-VkPhysicalDeviceExternalTensorInfoARM-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPhysicalDeviceExternalTensorInfoARM-flags-parameter)VUID-VkPhysicalDeviceExternalTensorInfoARM-flags-parameter  
  `flags` **must** be a valid combination of [VkTensorCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagBitsARM.html) values
- [](#VUID-VkPhysicalDeviceExternalTensorInfoARM-pDescription-parameter)VUID-VkPhysicalDeviceExternalTensorInfoARM-pDescription-parameter  
  `pDescription` **must** be a valid pointer to a valid [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) structure
- [](#VUID-VkPhysicalDeviceExternalTensorInfoARM-handleType-parameter)VUID-VkPhysicalDeviceExternalTensorInfoARM-handleType-parameter  
  `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateFlagsARM.html), [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html), [vkGetPhysicalDeviceExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalTensorPropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExternalTensorInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0