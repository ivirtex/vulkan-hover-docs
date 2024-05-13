# VkPhysicalDeviceMaintenance4Features(3) Manual Page

## Name

VkPhysicalDeviceMaintenance4Features - Structure describing whether the
implementation supports maintenance4 functionality



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance4Features` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceMaintenance4Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance4;
} VkPhysicalDeviceMaintenance4Features;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance4
typedef VkPhysicalDeviceMaintenance4Features VkPhysicalDeviceMaintenance4FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-maintenance4"></span> `maintenance4`
  indicates that the implementation supports the following:

  - The application **may** destroy a
    [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object immediately after
    using it to create another object.

  - `LocalSizeId` **can** be used as an alternative to `LocalSize` to
    specify the local workgroup size with specialization constants.

  - Images created with identical creation parameters will always have
    the same alignment requirements.

  - The size memory requirement of a buffer or image is never greater
    than that of another buffer or image created with a greater or equal
    size.

  - Push constants do not have to be initialized before they are
    dynamically accessed.

  - The interface matching rules allow a larger output vector to match
    with a smaller input vector, with additional values being discarded.

If the `VkPhysicalDeviceMaintenance4Features` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMaintenance4Features` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance4Features-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance4Features-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance4Features-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_4_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance4.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance4Features"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
