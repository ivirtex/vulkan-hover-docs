# VK_KHR_driver_properties(3) Manual Page

## Name

VK_KHR_driver_properties - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

197

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_driver_properties%5D%20@drakos-amd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_driver_properties%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>drakos-amd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-04-11

**IP Status**  
No known IP claims.

**Contributors**  
- Baldur Karlsson

- Matthaeus G. Chajdas, AMD

- Piers Daniell, NVIDIA

- Alexander Galazin, Arm

- Jesse Hall, Google

- Daniel Rakos, AMD

## <a href="#_description" class="anchor"></a>Description

This extension provides a new physical device query which allows
retrieving information about the driver implementation, allowing
applications to determine which physical device corresponds to which
particular vendor’s driver, and which conformance test suite version the
driver implementation is compliant with.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkConformanceVersionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConformanceVersionKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceDriverPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDriverPropertiesKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDriverIdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverIdKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DRIVER_PROPERTIES_EXTENSION_NAME`

- `VK_KHR_DRIVER_PROPERTIES_SPEC_VERSION`

- `VK_MAX_DRIVER_INFO_SIZE_KHR`

- `VK_MAX_DRIVER_NAME_SIZE_KHR`

- Extending [VkDriverId](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDriverId.html):

  - `VK_DRIVER_ID_AMD_OPEN_SOURCE_KHR`

  - `VK_DRIVER_ID_AMD_PROPRIETARY_KHR`

  - `VK_DRIVER_ID_ARM_PROPRIETARY_KHR`

  - `VK_DRIVER_ID_BROADCOM_PROPRIETARY_KHR`

  - `VK_DRIVER_ID_GGP_PROPRIETARY_KHR`

  - `VK_DRIVER_ID_GOOGLE_SWIFTSHADER_KHR`

  - `VK_DRIVER_ID_IMAGINATION_PROPRIETARY_KHR`

  - `VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR`

  - `VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR`

  - `VK_DRIVER_ID_MESA_RADV_KHR`

  - `VK_DRIVER_ID_NVIDIA_PROPRIETARY_KHR`

  - `VK_DRIVER_ID_QUALCOMM_PROPRIETARY_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-04-11 (Daniel Rakos)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_driver_properties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
