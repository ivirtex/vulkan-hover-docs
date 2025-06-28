# VK\_KHR\_variable\_pointers(3) Manual Page

## Name

VK\_KHR\_variable\_pointers - device extension



## [](#_registered_extension_number)Registered Extension Number

121

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK\_KHR\_storage\_buffer\_storage\_class](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_storage_buffer_storage_class.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_variable\_pointers](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_variable_pointers.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_variable_pointers%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_variable_pointers%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**IP Status**

No known IP claims.

**Contributors**

- John Kessenich, Google
- Neil Henning, Codeplay
- David Neto, Google
- Daniel Koch, Nvidia
- Graeme Leese, Broadcom
- Weifeng Zhang, Qualcomm
- Stephen Clarke, Imagination Technologies
- Faith Ekstrand, Intel
- Jesse Hall, Google

## [](#_description)Description

The `VK_KHR_variable_pointers` extension allows implementations to indicate their level of support for the `SPV_KHR_variable_pointers` SPIR-V extension. The SPIR-V extension allows shader modules to use invocation-private pointers into uniform and/or storage buffers, where the pointer values can be dynamic and non-uniform.

The `SPV_KHR_variable_pointers` extension introduces two capabilities. The first, `VariablePointersStorageBuffer`, **must** be supported by all implementations of this extension. The second, `VariablePointers`, is optional.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted, however support for the [`variablePointersStorageBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variablePointersStorageBuffer) feature is made optional. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

If Vulkan 1.4 is supported, support for the [`variablePointers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variablePointers) and [`variablePointersStorageBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variablePointersStorageBuffer) features is required.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVariablePointerFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVariablePointerFeaturesKHR.html)
  - [VkPhysicalDeviceVariablePointersFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVariablePointersFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_VARIABLE_POINTERS_EXTENSION_NAME`
- `VK_KHR_VARIABLE_POINTERS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTERS_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VARIABLE_POINTER_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`VariablePointers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-VariablePointers)
- [`VariablePointersStorageBuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-VariablePointersStorageBuffer)

## [](#_issues)Issues

1\) Do we need an optional property for the SPIR-V `VariablePointersStorageBuffer` capability or should it be mandatory when this extension is advertised?

**RESOLVED**: Add it as a distinct feature, but make support mandatory. Adding it as a feature makes the extension easier to include in a future core API version. In the extension, the feature is mandatory, so that presence of the extension guarantees some functionality. When included in a core API version, the feature would be optional.

2\) Can support for these capabilities vary between shader stages?

**RESOLVED**: No, if the capability is supported in any stage it must be supported in all stages.

3\) Should the capabilities be features or limits?

**RESOLVED**: Features, primarily for consistency with other similar extensions.

## [](#_version_history)Version History

- Revision 1, 2017-03-14 (Jesse Hall and John Kessenich)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_variable_pointers)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0