# VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT - Structure describing
whether the mutable descriptor type is supported



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_mutable_descriptor_type
typedef struct VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           mutableDescriptorType;
} VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT;
```

or the equivalent

``` c
// Provided by VK_VALVE_mutable_descriptor_type
typedef VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT VkPhysicalDeviceMutableDescriptorTypeFeaturesVALVE;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-mutableDescriptorType"></span>
  `mutableDescriptorType` indicates that the implementation **must**
  support using the [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) of
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` with at least the following
  descriptor types, where any combination of the types **must** be
  supported:

  - `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`

  - `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`

  - `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`

  - `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`

  - `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`

  - `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`

## <a href="#_description" class="anchor"></a>Description

- Additionally, `mutableDescriptorType` indicates that:

  - Non-uniform descriptor indexing **must** be supported if all
    descriptor types in a
    [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeListEXT.html)
    for `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` have the corresponding
    non-uniform indexing features enabled in
    [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html).

  - `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` with `descriptorType`
    of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` relaxes the list of required
    descriptor types to the descriptor types which have the
    corresponding update-after-bind feature enabled in
    [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html).

  - Dynamically uniform descriptor indexing **must** be supported if all
    descriptor types in a
    [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeListEXT.html)
    for `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` have the corresponding dynamic
    indexing features enabled.

  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT` **must** be
    supported.

  - `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` **must** be supported.

If the `VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MUTABLE_DESCRIPTOR_TYPE_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_mutable_descriptor_type](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mutable_descriptor_type.html),
[VK_VALVE_mutable_descriptor_type](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VALVE_mutable_descriptor_type.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
