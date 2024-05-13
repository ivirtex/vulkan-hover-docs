# VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR - Structure
describing the workgroup storage explicit layout features that can be
supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR` structure
is defined as:

``` c
// Provided by VK_KHR_workgroup_memory_explicit_layout
typedef struct VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           workgroupMemoryExplicitLayout;
    VkBool32           workgroupMemoryExplicitLayoutScalarBlockLayout;
    VkBool32           workgroupMemoryExplicitLayout8BitAccess;
    VkBool32           workgroupMemoryExplicitLayout16BitAccess;
} VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-workgroupMemoryExplicitLayout"></span>
  `workgroupMemoryExplicitLayout` indicates whether the implementation
  supports the SPIR-V `WorkgroupMemoryExplicitLayoutKHR` capability.

- <span id="features-workgroupMemoryExplicitLayoutScalarBlockLayout"></span>
  `workgroupMemoryExplicitLayoutScalarBlockLayout` indicates whether the
  implementation supports scalar alignment for laying out Workgroup
  Blocks.

- <span id="features-workgroupMemoryExplicitLayout8BitAccess"></span>
  `workgroupMemoryExplicitLayout8BitAccess` indicates whether objects in
  the `Workgroup` storage class with the `Block` decoration **can** have
  8-bit integer members. If this feature is not enabled, 8-bit integer
  members **must** not be used in such objects. This also indicates
  whether shader modules **can** declare the
  `WorkgroupMemoryExplicitLayout8BitAccessKHR` capability.

- <span id="features-workgroupMemoryExplicitLayout16BitAccess"></span>
  `workgroupMemoryExplicitLayout16BitAccess` indicates whether objects
  in the `Workgroup` storage class with the `Block` decoration **can**
  have 16-bit integer and 16-bit floating-point members. If this feature
  is not enabled, 16-bit integer or 16-bit floating-point members
  **must** not be used in such objects. This also indicates whether
  shader modules **can** declare the
  `WorkgroupMemoryExplicitLayout16BitAccessKHR` capability.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR` **can** also
be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_workgroup_memory_explicit_layout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_workgroup_memory_explicit_layout.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
