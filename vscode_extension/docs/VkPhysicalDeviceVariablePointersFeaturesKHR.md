# VkPhysicalDeviceVariablePointersFeatures(3) Manual Page

## Name

VkPhysicalDeviceVariablePointersFeatures - Structure describing variable
pointers features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVariablePointersFeatures` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceVariablePointersFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           variablePointersStorageBuffer;
    VkBool32           variablePointers;
} VkPhysicalDeviceVariablePointersFeatures;
```

``` c
// Provided by VK_VERSION_1_1
typedef VkPhysicalDeviceVariablePointersFeatures VkPhysicalDeviceVariablePointerFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_variable_pointers
typedef VkPhysicalDeviceVariablePointersFeatures VkPhysicalDeviceVariablePointersFeaturesKHR;
```

``` c
// Provided by VK_KHR_variable_pointers
typedef VkPhysicalDeviceVariablePointersFeatures VkPhysicalDeviceVariablePointerFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-variablePointersStorageBuffer"></span>
  `variablePointersStorageBuffer` specifies whether the implementation
  supports the SPIR-V `VariablePointersStorageBuffer` capability. When
  this feature is not enabled, shader modules **must** not declare the
  `SPV_KHR_variable_pointers` extension or the
  `VariablePointersStorageBuffer` capability.

- <span id="extension-features-variablePointers"></span>
  `variablePointers` specifies whether the implementation supports the
  SPIR-V `VariablePointers` capability. When this feature is not
  enabled, shader modules **must** not declare the `VariablePointers`
  capability.

If the `VkPhysicalDeviceVariablePointersFeatures` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceVariablePointersFeatures` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage

- <a
  href="#VUID-VkPhysicalDeviceVariablePointersFeatures-variablePointers-01431"
  id="VUID-VkPhysicalDeviceVariablePointersFeatures-variablePointers-01431"></a>
  VUID-VkPhysicalDeviceVariablePointersFeatures-variablePointers-01431  
  If `variablePointers` is enabled then `variablePointersStorageBuffer`
  **must** also be enabled

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceVariablePointersFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceVariablePointersFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceVariablePointersFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_variable_pointers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_variable_pointers.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVariablePointersFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
