# VK_KHR_8bit_storage(3) Manual Page

## Name

VK_KHR_8bit_storage - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

178

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
    
[VK_KHR_storage_buffer_storage_class](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_storage_buffer_storage_class.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_8bit_storage](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_8bit_storage.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Alexander Galazin <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_8bit_storage%5D%20@alegal-arm%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_8bit_storage%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alegal-arm</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-02-05

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_EXT_shader_16bit_storage`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shader_16bit_storage.txt)

**IP Status**  
No known IP claims.

**Contributors**  
- Alexander Galazin, Arm

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_8bit_storage` extension allows use of 8-bit types in uniform
and storage buffers, and push constant blocks. This extension introduces
several new optional features which map to SPIR-V capabilities and allow
access to 8-bit data in `Block`-decorated objects in the `Uniform` and
the `StorageBuffer` storage classes, and objects in the `PushConstant`
storage class.

The `StorageBuffer8BitAccess` capability **must** be supported by all
implementations of this extension. The other capabilities are optional.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

Functionality in this extension is included in core Vulkan 1.2, with the
KHR suffix omitted. However, if Vulkan 1.2 is supported and this
extension is not, the `StorageBuffer8BitAccess` capability is optional.
The original type, enum and command names are still available as aliases
of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevice8BitStorageFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice8BitStorageFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_8BIT_STORAGE_EXTENSION_NAME`

- `VK_KHR_8BIT_STORAGE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-StorageBuffer8BitAccess"
  target="_blank" rel="noopener"><code>StorageBuffer8BitAccess</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-UniformAndStorageBuffer8BitAccess"
  target="_blank"
  rel="noopener"><code>UniformAndStorageBuffer8BitAccess</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-StoragePushConstant8"
  target="_blank" rel="noopener"><code>StoragePushConstant8</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-02-05 (Alexander Galazin)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_8bit_storage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
