# VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT - Structure
describing whether list type primitives can support primitive restart



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT` structure
is defined as:

``` c
// Provided by VK_EXT_primitive_topology_list_restart
typedef struct VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           primitiveTopologyListRestart;
    VkBool32           primitiveTopologyPatchListRestart;
} VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-primitiveTopologyListRestart"></span>
  `primitiveTopologyListRestart` indicates that list type primitives,
  `VK_PRIMITIVE_TOPOLOGY_POINT_LIST`, `VK_PRIMITIVE_TOPOLOGY_LINE_LIST`,
  `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST`,
  `VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY` and
  `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY`, **can** use the
  primitive restart index value in index buffers.

- <span id="features-primitiveTopologyPatchListRestart"></span>
  `primitiveTopologyPatchListRestart` indicates that the
  `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST` topology **can** use the primitive
  restart index value in index buffers.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT` **can** also
be used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIMITIVE_TOPOLOGY_LIST_RESTART_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_primitive_topology_list_restart](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_primitive_topology_list_restart.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
