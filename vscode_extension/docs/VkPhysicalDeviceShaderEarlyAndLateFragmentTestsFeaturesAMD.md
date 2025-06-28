# VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD(3) Manual Page

## Name

VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD - Structure describing whether early and late fragment tests can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD` structure is defined as:

```c++
// Provided by VK_AMD_shader_early_and_late_fragment_tests
typedef struct VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderEarlyAndLateFragmentTests;
} VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`shaderEarlyAndLateFragmentTests` indicates whether the implementation supports the `EarlyAndLateFragmentTestsAMD` `Execution` `Mode`.

## [](#_description)Description

If the `VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD-sType-sType)VUID-VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_EARLY_AND_LATE_FRAGMENT_TESTS_FEATURES_AMD`

## [](#_see_also)See Also

[VK\_AMD\_shader\_early\_and\_late\_fragment\_tests](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_shader_early_and_late_fragment_tests.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0