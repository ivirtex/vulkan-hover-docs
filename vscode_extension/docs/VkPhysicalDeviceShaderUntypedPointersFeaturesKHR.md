# VkPhysicalDeviceShaderUntypedPointersFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceShaderUntypedPointersFeaturesKHR - Structure describing support for untyped pointers in shader by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderUntypedPointersFeaturesKHR` structure is defined as:

```c++
// Provided by VK_KHR_shader_untyped_pointers
typedef struct VkPhysicalDeviceShaderUntypedPointersFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderUntypedPointers;
} VkPhysicalDeviceShaderUntypedPointersFeaturesKHR;
```

## [](#_members)Members

The members of `VkPhysicalDeviceShaderUntypedPointersFeaturesKHR` describe the following features:

## [](#_description)Description

- []()`shaderUntypedPointers` specifies whether shader modules **can** declare the `UntypedPointersKHR` capability and untyped pointers in any [explicitly laid out storage class](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-resources-layout).

If the `VkPhysicalDeviceShaderUntypedPointersFeaturesKHR` structure is included in the `pNext` chain of [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), it is filled with values indicating whether the features are supported. `VkPhysicalDeviceShaderUntypedPointersFeaturesKHR` **can** also be included in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) to enable the features.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderUntypedPointersFeaturesKHR-sType-sType)VUID-VkPhysicalDeviceShaderUntypedPointersFeaturesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_UNTYPED_POINTERS_FEATURES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_shader\_untyped\_pointers](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_untyped_pointers.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderUntypedPointersFeaturesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0