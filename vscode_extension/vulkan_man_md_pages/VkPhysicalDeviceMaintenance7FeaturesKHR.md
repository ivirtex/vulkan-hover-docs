# VkPhysicalDeviceMaintenance7FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance7FeaturesKHR - Structure describing whether
the implementation supports maintenance7 functionality



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance7FeaturesKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceMaintenance7FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance7;
} VkPhysicalDeviceMaintenance7FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-maintenance7"></span> `maintenance7` indicates that
  the implementation supports the following:

  - The `VK_RENDERING_CONTENTS_INLINE_BIT_KHR` and
    `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR` flags
    **can** be used to record commands in render pass instances both
    inline and in secondary command buffers executed with
    [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html) for dynamic
    rendering and legacy render passes respectively.

  - Querying information regarding the underlying devices in
    environments where the Vulkan implementation is provided through
    layered implementations. This is done by chaining
    [VkPhysicalDeviceLayeredApiPropertiesListKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesListKHR.html)
    to [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html).

  - New limits which indicate the maximum total count of dynamic uniform
    buffers and dynamic storage buffers that **can** be included in a
    pipeline layout.

  - 32-bit timestamp queries **must** wrap on overflow

  - A property that indicates whether a fragment shading rate attachment
    can have a size that is too small to cover a specified render area.

  - A property that indicates support for writing to one aspect of a
    depth/stencil attachment without performing a read-modify-write
    operation on the other aspect

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMaintenance7FeaturesKHR` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMaintenance7FeaturesKHR` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance7FeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance7FeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance7FeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance7](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance7.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance7FeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
