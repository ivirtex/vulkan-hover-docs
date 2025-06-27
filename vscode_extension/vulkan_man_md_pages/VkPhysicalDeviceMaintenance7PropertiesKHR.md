# VkPhysicalDeviceMaintenance7PropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance7PropertiesKHR - Structure describing various
implementation-defined properties introduced with VK_KHR_maintenance7



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance7PropertiesKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-robustFragmentShadingRateAttachmentAccess"></span>
  `robustFragmentShadingRateAttachmentAccess` indicates whether the
  scaled size of a fragment shading rate attachment **can** be less than
  the size of the render area. If
  `robustFragmentShadingRateAttachmentAccess` is `VK_FALSE`, the size of
  the attachment multiplied by the texel size **must** be greater than
  or equal to the size of the render area. If it is `VK_TRUE` and the
  fragment shading rate attachment was created with
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseMipLevel`
  equal to 0, the scaled size **can** be smaller than the render area,
  and shading rates for missing texels are defined by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-texel-replacement"
  target="_blank" rel="noopener">texel replacement for invalid texels</a>.

- <span id="limits-separateDepthStencilAttachmentAccess"></span>
  `separateDepthStencilAttachmentAccess` indicates support for writing
  to one aspect of a depth/stencil attachment without performing
  read-modify-write operations on the other aspect. If this property is
  `VK_TRUE`, writes to one aspect **must** not result in
  read-modify-write operations on the other aspect. If `VK_FALSE`,
  writes to one aspect **may** result in writes to the other aspect as
  defined by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
  target="_blank" rel="noopener">render pass load operations</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
  target="_blank" rel="noopener">render pass store operations</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
  target="_blank" rel="noopener">render pass resolve operations</a>.

- <span id="limits-maxDescriptorSetTotalUniformBuffersDynamic"></span>
  `maxDescriptorSetTotalUniformBuffersDynamic` is the maximum total
  count of dynamic uniform buffers that **can** be included in a
  pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic</a>.

- <span id="limits-maxDescriptorSetTotalStorageBuffersDynamic"></span>
  `maxDescriptorSetTotalStorageBuffersDynamic` is the maximum total
  count of dynamic storage buffers that **can** be included in a
  pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic</a>.

- <span id="limits-maxDescriptorSetTotalBuffersDynamic"></span>
  `maxDescriptorSetTotalBuffersDynamic` is the maximum total count of
  dynamic uniform buffers and storage buffers that **can** be included
  in a pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit.

- <span id="limits-maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic"></span>
  `maxDescriptorSetUpdateAfterBindTotalUniformBuffersDynamic` is similar
  to `maxDescriptorSetTotalUniformBuffersDynamic` but counts descriptors
  from descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="limits-maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic"></span>
  `maxDescriptorSetUpdateAfterBindTotalStorageBuffersDynamic` is similar
  to `maxDescriptorSetTotalStorageBuffersDynamic` but counts descriptors
  from descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="limits-maxDescriptorSetUpdateAfterBindTotalBuffersDynamic"></span>
  `maxDescriptorSetUpdateAfterBindTotalBuffersDynamic` is similar to
  `maxDescriptorSetTotalBuffersDynamic` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
  While an application **can** allocate dynamic storage buffer
  descriptors from a pool created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, bindings
  for these descriptors **must** not be present in any descriptor set
  layout that includes bindings created with
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMaintenance7PropertiesKHR` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance7PropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance7PropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance7PropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance7](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance7.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance7PropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
