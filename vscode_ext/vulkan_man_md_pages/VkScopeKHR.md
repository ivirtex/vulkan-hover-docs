# VkScopeKHR(3) Manual Page

## Name

VkScopeKHR - Specify SPIR-V scope



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values for [VkScopeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeKHR.html) include:

``` c
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

``` c
// Provided by VK_NV_cooperative_matrix
typedef VkScopeKHR VkScopeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SCOPE_DEVICE_KHR` corresponds to SPIR-V `Device` scope.

- `VK_SCOPE_WORKGROUP_KHR` corresponds to SPIR-V `Workgroup` scope.

- `VK_SCOPE_SUBGROUP_KHR` corresponds to SPIR-V `Subgroup` scope.

- `VK_SCOPE_QUEUE_FAMILY_KHR` corresponds to SPIR-V `QueueFamily` scope.

All enum values match the corresponding SPIR-V value.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_cooperative_matrix](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_cooperative_matrix.html),
[VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkScopeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
