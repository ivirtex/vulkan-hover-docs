# VkPhysicalDeviceMaintenance7PropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance7PropertiesKHR - Structure describing various implementation-defined properties introduced with VK\_KHR\_maintenance7



## [](#_c_specification)C Specification

The `VkPhysicalDeviceMaintenance7PropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_maintenance7
typedef struct VkPhysicalDeviceMaintenance7PropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           robustFragmentShadingRateAttachmentAccess;
    VkBool32           separateDepthStencilAttachmentAccess;
    uint32_t           maxDescriptorSetTotalUniformBuffersDynamic;
    uint32_t           maxDescriptorSetTotalStorageBuffersDynamic;
    uint32_t           maxDescriptorSetTotalBuffersDynamic;
    uint32_t           maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic;
    uint32_t           maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic;
    uint32_t           maxDescriptorSetUpdateAfterBindTotalBuffersDynamic;
} VkPhysicalDeviceMaintenance7PropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`robustFragmentShadingRateAttachmentAccess` indicates whether the scaled size of a fragment shading rate attachment **can** be less than the size of the render area. If `robustFragmentShadingRateAttachmentAccess` is `VK_FALSE`, the size of the attachment multiplied by the texel size **must** be greater than or equal to the size of the render area. If it is `VK_TRUE` and the fragment shading rate attachment was created with [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseMipLevel` equal to 0, the scaled size **can** be smaller than the render area, and shading rates for missing texels are defined by [texel replacement for invalid texels](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-texel-replacement).
- []()`separateDepthStencilAttachmentAccess` indicates support for writing to one aspect of a depth/stencil attachment without performing read-modify-write operations on the other aspect. If this property is `VK_TRUE`, writes to one aspect **must** not result in read-modify-write operations on the other aspect. If `VK_FALSE`, writes to one aspect **may** result in writes to the other aspect as defined by [render pass load operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations), [render pass store operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations) and [render pass resolve operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations).
- []()`maxDescriptorSetTotalUniformBuffersDynamic` is the maximum total count of dynamic uniform buffers that **can** be included in a pipeline layout. Descriptors with a type of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` count against this limit. Only descriptors in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit. See [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformbufferdynamic](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformbufferdynamic).
- []()`maxDescriptorSetTotalStorageBuffersDynamic` is the maximum total count of dynamic storage buffers that **can** be included in a pipeline layout. Descriptors with a type of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` count against this limit. Only descriptors in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit. See [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebufferdynamic](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebufferdynamic).
- []()`maxDescriptorSetTotalBuffersDynamic` is the maximum total count of dynamic uniform buffers and storage buffers that **can** be included in a pipeline layout. Descriptors with a type of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` count against this limit. Only descriptors in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit.
- []()`maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic` is similar to `maxDescriptorSetTotalUniformBuffersDynamic` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic` is similar to `maxDescriptorSetTotalStorageBuffersDynamic` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindTotalBuffersDynamic` is similar to `maxDescriptorSetTotalBuffersDynamic` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set. While an application **can** allocate dynamic storage buffer descriptors from a pool created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, bindings for these descriptors **must** not be present in any descriptor set layout that includes bindings created with `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`.

## [](#_description)Description

If the `VkPhysicalDeviceMaintenance7PropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMaintenance7PropertiesKHR-sType-sType)VUID-VkPhysicalDeviceMaintenance7PropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_maintenance7](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance7.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMaintenance7PropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0