# VkScopeKHR(3) Manual Page

## Name

VkScopeKHR - Specify SPIR-V scope



## [](#_c_specification)C Specification

Possible values for [VkScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScopeKHR.html) include:

```c++
// Provided by VK_KHR_cooperative_matrix
typedef enum VkScopeKHR {
    VK_SCOPE_DEVICE_KHR = 1,
    VK_SCOPE_WORKGROUP_KHR = 2,
    VK_SCOPE_SUBGROUP_KHR = 3,
    VK_SCOPE_QUEUE_FAMILY_KHR = 5,
  // Provided by VK_NV_cooperative_matrix
    VK_SCOPE_DEVICE_NV = VK_SCOPE_DEVICE_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_SCOPE_WORKGROUP_NV = VK_SCOPE_WORKGROUP_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_SCOPE_SUBGROUP_NV = VK_SCOPE_SUBGROUP_KHR,
  // Provided by VK_NV_cooperative_matrix
    VK_SCOPE_QUEUE_FAMILY_NV = VK_SCOPE_QUEUE_FAMILY_KHR,
} VkScopeKHR;
```

or the equivalent

```c++
// Provided by VK_NV_cooperative_matrix
typedef VkScopeKHR VkScopeNV;
```

## [](#_description)Description

- `VK_SCOPE_DEVICE_KHR` corresponds to SPIR-V `Device` scope.
- `VK_SCOPE_WORKGROUP_KHR` corresponds to SPIR-V `Workgroup` scope.
- `VK_SCOPE_SUBGROUP_KHR` corresponds to SPIR-V `Subgroup` scope.
- `VK_SCOPE_QUEUE_FAMILY_KHR` corresponds to SPIR-V `QueueFamily` scope.

All enum values match the corresponding SPIR-V value.

## [](#_see_also)See Also

[VK\_KHR\_cooperative\_matrix](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_cooperative_matrix.html), [VkCooperativeMatrixFlexibleDimensionsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeMatrixFlexibleDimensionsPropertiesNV.html), [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeMatrixPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkScopeKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0