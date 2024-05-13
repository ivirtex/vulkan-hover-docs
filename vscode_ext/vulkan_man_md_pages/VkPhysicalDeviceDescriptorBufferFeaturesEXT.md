# VkPhysicalDeviceDescriptorBufferFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceDescriptorBufferFeaturesEXT - Structure describing the
descriptor buffer features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDescriptorBufferFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkPhysicalDeviceDescriptorBufferFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           descriptorBuffer;
    VkBool32           descriptorBufferCaptureReplay;
    VkBool32           descriptorBufferImageLayoutIgnored;
    VkBool32           descriptorBufferPushDescriptors;
} VkPhysicalDeviceDescriptorBufferFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-descriptorBuffer"></span> `descriptorBuffer`
  indicates that the implementation supports putting shader-accessible
  descriptors directly in memory.

- <span id="features-descriptorBufferCaptureReplay"></span>
  `descriptorBufferCaptureReplay` indicates that the implementation
  supports capture and replay when using descriptor buffers. If this is
  `VK_TRUE`, all resources created with
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`,
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`,
  `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`,
  `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`, or
  `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
  **must** be created before resources of the same types without those
  flags.

- <span id="features-descriptorBufferImageLayoutIgnored"></span>
  `descriptorBufferImageLayoutIgnored` indicates that the implementation
  will ignore `imageLayout` in `VkDescriptorImageInfo` when calling
  [vkGetDescriptorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorEXT.html).

- <span id="features-descriptorBufferPushDescriptors"></span>
  `descriptorBufferPushDescriptors` indicates that the implementation
  supports using push descriptors with descriptor buffers.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceDescriptorBufferFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceDescriptorBufferFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceDescriptorBufferFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceDescriptorBufferFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceDescriptorBufferFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDescriptorBufferFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
