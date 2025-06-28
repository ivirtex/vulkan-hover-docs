# VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT - Structure describing whether the mutable descriptor type is supported



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_mutable_descriptor_type
typedef struct VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           mutableDescriptorType;
} VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT;
```

or the equivalent

```c++
// Provided by VK_VALVE_mutable_descriptor_type
typedef VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT VkPhysicalDeviceMutableDescriptorTypeFeaturesVALVE;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`mutableDescriptorType` indicates that the implementation **must** support using the [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` with at least the following descriptor types, where any combination of the types **must** be supported:
  
  - `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`
  - `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`
  - `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`
  - `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`
  - `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`
  - `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`

## [](#_description)Description

- Additionally, `mutableDescriptorType` indicates that:
  
  - Non-uniform descriptor indexing **must** be supported if all descriptor types in a [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeListEXT.html) for `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` have the corresponding non-uniform indexing features enabled in [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html).
  - `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` with `descriptorType` of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` relaxes the list of required descriptor types to the descriptor types which have the corresponding update-after-bind feature enabled in [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html).
  - Dynamically uniform descriptor indexing **must** be supported if all descriptor types in a [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeListEXT.html) for `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` have the corresponding dynamic indexing features enabled.
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT` **must** be supported.
  - `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` **must** be supported.

If the `VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MUTABLE_DESCRIPTOR_TYPE_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_mutable\_descriptor\_type](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_mutable_descriptor_type.html), [VK\_VALVE\_mutable\_descriptor\_type](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_mutable_descriptor_type.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0